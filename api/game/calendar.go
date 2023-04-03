package game

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"gitlab.com/khakibee/khakibee/api/logger"
	"gitlab.com/khakibee/khakibee/api/model"
	"gitlab.com/khakibee/khakibee/api/user"
)

func (g *GameHandler) GetCalendarSSE(c echo.Context) error {
	var (
		gm     model.Game
		offset int
		loc    time.Location
	)
	req := new(getCalendarReq)

	if err := req.bind(c, &gm, &offset, &loc); err != nil {
		return logger.EchoHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	userID := c.Get(user.UserIDCtxKey)
	openBkngsOnly := userID == user.UserIDLimitedAccess

	clndr, echoErr := getCalendar(c.Request().Context(), g, gm.GameID, loc, offset, openBkngsOnly)
	if echoErr != nil {
		return echoErr
	}

	c.Response().Header().Set("Content-Type", "text/event-stream")
	c.Response().Header().Set("Cache-Control", "no-cache")
	c.Response().Header().Set("Connection", "keep-alive")

	clndrBt, err := json.Marshal(clndr)
	if err != nil {
		return logger.EchoHTTPError(http.StatusInternalServerError, err)
	}

	fmt.Fprintf(c.Response().Writer, "data: %v\n\n", string(clndrBt))
	c.Response().Flush()

	for {
		select {
		case <-g.bookingChanged:
			clndr, echoErr := getCalendar(c.Request().Context(), g, gm.GameID, loc, offset, openBkngsOnly)
			if echoErr != nil {
				return echoErr
			}

			clndrBt, err := json.Marshal(clndr)
			if err != nil {
				return logger.EchoHTTPError(http.StatusInternalServerError, err)
			}

			fmt.Fprintf(c.Response().Writer, "data: %v\n\n", string(clndrBt))
			c.Response().Flush()
		case <-c.Request().Context().Done():
			return nil
		}
	}

}

func (g *GameHandler) GetCalendar(c echo.Context) error {
	var (
		gm     model.Game
		offset int
		loc    time.Location
	)
	req := new(getCalendarReq)

	if err := req.bind(c, &gm, &offset, &loc); err != nil {
		return logger.EchoHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	userID := c.Get(user.UserIDCtxKey)
	openBkngsOnly := userID == user.UserIDLimitedAccess

	clndr, err := getCalendar(c.Request().Context(), g, gm.GameID, loc, offset, openBkngsOnly)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, clndr)
}

func getCalendar(ctx context.Context, g *GameHandler, gmeID int, loc time.Location, offset int, openBkngsOnly bool) (calendar []model.Event, echoErr *echo.HTTPError) {
	if openBkngsOnly && offset != 0 && offset != 1 {
		return []model.Event{}, echoErr
	}

	schdl, err := g.schdlStore.SelectScheduleByGameID(ctx, gmeID, false)
	if err != nil {
		echoErr = logger.EchoHTTPError(http.StatusInternalServerError, err)
		return
	}

	schdlUpcoming, err := g.schdlStore.SelectScheduleByGameID(ctx, gmeID, true)
	if err != nil {
		echoErr = logger.EchoHTTPError(http.StatusInternalServerError, err)
		return
	}

	var bkngs []model.Booking

	bkngs, err = g.bkngStore.SelectBookings(ctx, gmeID, loc.String(), offset)
	if err != nil {
		echoErr = logger.EchoHTTPError(http.StatusInternalServerError, err)
		return
	}

	calendar, err = CreateCalendar(schdl, schdlUpcoming, bkngs, offset, &loc, openBkngsOnly)
	if err != nil {
		echoErr = logger.EchoHTTPError(http.StatusInternalServerError, err)
	}

	return
}

func CreateCalendar(schdl *model.Schedule, schdlUpcoming *model.Schedule, bkngs []model.Booking, offset int, loc *time.Location, openBkngsOnly bool) (calendar []model.Event, err error) {
	c := []model.Event{}

	// get date by offset on specific location
	t := time.Now().In(loc).AddDate(0, offset, 0)
	_, timezoneOffset := t.Zone()

	year, month, day := t.Date()

	//get last date of month
	_, _, lastDay := (t.AddDate(0, 1, -day)).Date()

	firstDayOfCalendar := 1
	if offset == 0 {
		firstDayOfCalendar = day
	}

	// get activation date of current schedule if it exists
	var schdlActvDte *time.Time
	if schdl != nil {
		schdlActvDte = &time.Time{}
		*schdlActvDte, err = time.Parse("2006-01-02", schdl.ActiveBy)
		if err != nil {
			return nil, err
		}
		// add timezone and timezone offset for correct comparisons with current date
		*schdlActvDte = schdlActvDte.Add(time.Duration(-timezoneOffset) * time.Second).In(loc)
	}

	// get activation date of upcoming schedule if it exists
	var schdlUpcActvDte *time.Time
	if schdlUpcoming != nil {
		schdlUpcActvDte = &time.Time{}
		*schdlUpcActvDte, err = time.Parse("2006-01-02", schdlUpcoming.ActiveBy)
		if err != nil {
			return nil, err
		}
		// add timezone and timezone offset for correct comparisons with current date
		*schdlUpcActvDte = schdlUpcActvDte.Add(time.Duration(-timezoneOffset) * time.Second).In(loc)
	}

	//generate events according to schedule
	for currMonthDay := firstDayOfCalendar; currMonthDay <= lastDay && schdl != nil; currMonthDay++ {
		wkdayTm := time.Date(year, month, currMonthDay, 0, 0, 0, 0, t.Location())
		wkday := wkdayTm.Weekday().String()

		var slots []string
		if schdlUpcActvDte != nil && !wkdayTm.Before(*schdlUpcActvDte) {
			slots = GetWeekdaySlots(schdlUpcoming.WkSchdl, wkday)
		} else if !wkdayTm.Before(*schdlActvDte) {
			slots = GetWeekdaySlots(schdl.WkSchdl, wkday)
		}

		for _, slot := range slots {
			s := strings.Split(slot, ":")

			var hours, minutes int
			if hours, err = strconv.Atoi(s[0]); err != nil {
				return nil, err
			}
			if minutes, err = strconv.Atoi(s[1]); err != nil {
				return nil, err
			}

			eventStartTm := time.Date(year, month, currMonthDay, hours, minutes, 0, 0, t.Location())

			eventStart := eventStartTm.Format("2006-01-02 15:04")
			eventEnd := eventStartTm.Add(time.Minute * time.Duration(schdl.Dur)).Format("2006-01-02 15:04")

			event := model.Event{
				Start: eventStart,
				End:   eventEnd,
			}

			c = append(c, event)
		}
	}

	//generete events from bookings
	for _, b := range bkngs {
		var bkngDte time.Time
		var inSchedule bool
		if bkngDte, err = time.Parse("2006-01-02 15:04", b.Dte); err != nil {
			return nil, err
		}
		bkngDte = bkngDte.In(loc)

		if b.BkngID == 106 {
			fmt.Printf("booking %v\n\n", bkngDte)

		}

		eventStart := bkngDte.Format("2006-01-02 15:04")
		eventEnd := bkngDte.Add(time.Minute * time.Duration(schdl.Dur)).Format("2006-01-02 15:04")

		event := model.Event{
			Start:   eventStart,
			End:     eventEnd,
			Booking: &model.Booking{BkngID: b.BkngID, Name: b.Name},
		}

		// if the booking event is on schedule, update calendar.
		for idx, e := range c {
			if e.Start == event.Start {
				if !openBkngsOnly { // add booking to calendar
					c[idx] = event
					inSchedule = true
					break
				} else { // if calendar is with limited view, remove it from calendarr
					c[idx] = c[len(c)-1]
					c = c[:len(c)-1]
				}
			}
		}

		// if calendar is not limited and booking is not on schedule, mark it and add it on callendar
		if !inSchedule && !openBkngsOnly {
			if !bkngDte.Before(t) {
				event.OutOfSchedule = true
			}
			c = append(c, event)
		}
	}

	sort.Slice(c, func(i, j int) bool {
		return c[i].Start < c[j].Start
	})

	return c, nil
}

func GetWeekdaySlots(ws model.WeekSchedule, weekday string) (slots []string) {
	wsValues := reflect.ValueOf(ws)
	slots = reflect.Indirect(wsValues).FieldByName(weekday).Interface().([]string)
	return
}
