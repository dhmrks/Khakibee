package game_test

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"gitlab.com/khakibee/khakibee/api/game"
	"gitlab.com/khakibee/khakibee/api/model"
	"gitlab.com/khakibee/khakibee/api/testUtils"
	"gitlab.com/khakibee/khakibee/api/user"
)

func TestGetCalendarSSE(t *testing.T) {
	setup()
	gid := strconv.Itoa(testUtils.DB1.Games[0].GameID)

	loc, _ := time.LoadLocation(testUtils.Timezone)

	bs.Bookings = []model.Booking{}

	schdl := testUtils.DB1.Schedule[1][0]
	schdl2 := testUtils.DB1.Schedule[1][1]
	wantCalendar, err := game.CreateCalendar(schdl, schdl2, []model.Booking{}, 0, loc, false)
	testUtils.AssertError(t, err, false)

	uid := testUtils.DB1.Users[0].UserID

	t.Run("get calendar sse successfully", func(t *testing.T) {
		c, _, res := testUtils.NewGetRequest()
		c.SetParamNames("gid")
		c.SetParamValues(gid)

		c.Set(user.UserIDCtxKey, uid)

		go func() {
			err := h.GetCalendarSSE(c)
			testUtils.AssertError(t, err, false)
		}()

		testUtils.AssertEqual(t, res.Flushed, false)
		time.Sleep(time.Millisecond) // wait 1 millisecond for the first data to be flushed
		testUtils.AssertEqual(t, res.Flushed, true)

		//get calendar when estabilishing connection
		bd := res.Body.String()
		bd = strings.Replace(bd, "data: ", "", 1)

		var gotCalendar []model.Event
		err := json.Unmarshal([]byte(bd), &gotCalendar)
		testUtils.AssertError(t, err, false)

		testUtils.AssertHTTPStatus(t, res.Code, http.StatusOK)
		testUtils.AssertEqual(t, gotCalendar, wantCalendar)

		h.DispatchBookingUpdate()    // dispach booking update signal
		time.Sleep(time.Millisecond) // wait 1 millisecond for the data to be flushed

		//get calendar when a booking is being updated
		bd = res.Body.String()
		bd = strings.Replace(bd, "data: ", "[", 1)
		bd = strings.Replace(bd, "data: ", ",", 1) + "]"

		var gotCalendars [][]model.Event
		err = json.Unmarshal([]byte(bd), &gotCalendars)
		testUtils.AssertError(t, err, false)

		wantCalendars := [][]model.Event{
			wantCalendar, wantCalendar,
		}

		testUtils.AssertHTTPStatus(t, res.Code, http.StatusOK)
		testUtils.AssertEqual(t, gotCalendars, wantCalendars)
	})

	setup() // re initializing handler to reset channel

	bs.Bookings = []model.Booking{}

	wantCalendarLimited, err := game.CreateCalendar(schdl, schdl2, []model.Booking{}, 0, loc, true)
	testUtils.AssertError(t, err, false)

	t.Run("get limited access calendar sse successfully", func(t *testing.T) {
		c, _, res := testUtils.NewGetRequest()
		c.SetParamNames("gid")
		c.SetParamValues(gid)

		c.Set(user.UserIDCtxKey, user.UserIDLimitedAccess)

		go func() {
			err := h.GetCalendarSSE(c)
			testUtils.AssertError(t, err, false)
		}()

		testUtils.AssertEqual(t, res.Flushed, false)
		time.Sleep(time.Millisecond) // wait 1 millisecond for the first data to be flushed
		testUtils.AssertEqual(t, res.Flushed, true)

		//get calendar when estabilishing connection
		bd := res.Body.String()
		bd = strings.Replace(bd, "data: ", "", 1)

		var gotCalendar []model.Event
		err := json.Unmarshal([]byte(bd), &gotCalendar)
		testUtils.AssertError(t, err, false)

		testUtils.AssertHTTPStatus(t, res.Code, http.StatusOK)
		testUtils.AssertEqual(t, gotCalendar, wantCalendarLimited)

		h.DispatchBookingUpdate()    // dispach booking update signal
		time.Sleep(time.Millisecond) // wait 1 millisecond for the data to be flushed

		//get calendar when a booking is being updated
		bd = res.Body.String()
		bd = strings.Replace(bd, "data: ", "[", 1)
		bd = strings.Replace(bd, "data: ", ",", 1) + "]"

		var gotCalendars [][]model.Event
		err = json.Unmarshal([]byte(bd), &gotCalendars)
		testUtils.AssertError(t, err, false)

		wantCalendars := [][]model.Event{
			wantCalendarLimited, wantCalendarLimited,
		}

		testUtils.AssertHTTPStatus(t, res.Code, http.StatusOK)
		testUtils.AssertEqual(t, gotCalendars, wantCalendars)
	})

	t.Run("get calendar sse with error", func(t *testing.T) {
		c, _, _ := testUtils.NewGetRequest()
		c.SetParamNames("gid")
		c.SetParamValues(gid)

		c.Set(user.UserIDCtxKey, uid)

		bs.ReturnError = errStore

		err := h.GetCalendar(c).(*echo.HTTPError)

		testUtils.AssertError(t, err, true)
		testUtils.AssertHTTPStatus(t, err.Code, http.StatusInternalServerError)
	})

}

func TestGetCalendar(t *testing.T) {
	setup()
	gid := strconv.Itoa(testUtils.DB1.Games[0].GameID)
	gid2 := strconv.Itoa(testUtils.DB1.Games[2].GameID)

	loc, _ := time.LoadLocation(testUtils.Timezone)

	bs.Bookings = []model.Booking{}

	schdl := testUtils.DB1.Schedule[1][0]
	schdl2 := testUtils.DB1.Schedule[1][1]
	wantCalendar, err := game.CreateCalendar(schdl, schdl2, []model.Booking{}, 0, loc, false)
	testUtils.AssertError(t, err, false)

	wantCalendarLimited, err := game.CreateCalendar(schdl, schdl2, []model.Booking{}, 0, loc, true)
	testUtils.AssertError(t, err, false)

	uid := testUtils.DB1.Users[0].UserID

	t.Run("get calendar successfully", func(t *testing.T) {

		c, _, res := testUtils.NewGetRequest()
		c.SetParamNames("gid")
		c.SetParamValues(gid)

		c.Set(user.UserIDCtxKey, uid)

		err := h.GetCalendar(c)

		testUtils.AssertError(t, err, false)

		var gotCalendar []model.Event
		json.NewDecoder(res.Body).Decode(&gotCalendar)

		testUtils.AssertHTTPStatus(t, res.Code, http.StatusOK)
		testUtils.AssertEqual(t, gotCalendar, wantCalendar)
	})

	t.Run("get calendar limited access successfully", func(t *testing.T) {

		c, _, res := testUtils.NewGetRequest()
		c.SetParamNames("gid")
		c.SetParamValues(gid)

		c.Set(user.UserIDCtxKey, user.UserIDLimitedAccess)

		err := h.GetCalendar(c)

		testUtils.AssertError(t, err, false)

		var gotCalendar []model.Event
		json.NewDecoder(res.Body).Decode(&gotCalendar)

		testUtils.AssertHTTPStatus(t, res.Code, http.StatusOK)
		testUtils.AssertEqual(t, gotCalendar, wantCalendarLimited)
	})

	t.Run("get calendar with empty schedule", func(t *testing.T) {
		c, _, res := testUtils.NewGetRequest()
		c.SetParamNames("gid")
		c.SetParamValues(gid2)

		err := h.GetCalendar(c)

		testUtils.AssertError(t, err, false)

		var gotCalendar []model.Event
		json.NewDecoder(res.Body).Decode(&gotCalendar)

		testUtils.AssertHTTPStatus(t, res.Code, http.StatusOK)
		testUtils.AssertEqual(t, len(gotCalendar), 0)
	})

	t.Run("get calendar with invalid game id", func(t *testing.T) {
		c, _, _ := testUtils.NewGetRequest()
		c.SetParamNames("gid")
		c.SetParamValues("a")

		err := h.GetCalendar(c).(*echo.HTTPError)

		testUtils.AssertError(t, err, true)
		testUtils.AssertHTTPStatus(t, err.Code, http.StatusUnprocessableEntity)
	})

	t.Run("get calendar with store error", func(t *testing.T) {
		c, _, _ := testUtils.NewGetRequest()
		c.SetParamNames("gid")
		c.SetParamValues(gid)

		c.Set(user.UserIDCtxKey, uid)

		bs.ReturnError = errStore

		err := h.GetCalendar(c).(*echo.HTTPError)

		testUtils.AssertError(t, err, true)
		testUtils.AssertHTTPStatus(t, err.Code, http.StatusInternalServerError)
	})

}
