package game

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gitlab.com/khakibee/khakibee/api/logger"
	"gitlab.com/khakibee/khakibee/api/model"
)

func (g GameHandler) PutNewSchedule(c echo.Context) error {
	var (
		gmeID int
		s     model.WeekSchedule
		ab    string
	)
	req := new(putNewScheduleReq)

	if err := req.bind(c, &gmeID, &s, &ab); err != nil {
		return logger.EchoHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	upcomingSchdl, err := g.schdlStore.SelectScheduleByGameID(c.Request().Context(), gmeID, true)
	if err != nil {
		return logger.EchoHTTPError(http.StatusInternalServerError, err)
	}

	var schdlID *int
	if upcomingSchdl != nil {
		schdlID = &upcomingSchdl.SchdlID
	}

	g.schdlStore.UpsertNewSchedule(c.Request().Context(), gmeID, s, ab, schdlID)
	if err != nil {
		return logger.EchoHTTPError(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusNoContent)
}

func (g GameHandler) PutSchedule(c echo.Context) error {
	var (
		gmeID int
		s     model.WeekSchedule
	)
	req := new(putScheduleReq)

	if err := req.bind(c, &gmeID, &s); err != nil {
		return logger.EchoHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	currentSchdl, err := g.schdlStore.SelectScheduleByGameID(c.Request().Context(), gmeID, false)
	if err != nil {
		return logger.EchoHTTPError(http.StatusInternalServerError, err)
	}

	var schdlID *int
	if currentSchdl != nil {
		schdlID = &currentSchdl.SchdlID
	}

	err = g.schdlStore.UpsertSchedule(c.Request().Context(), gmeID, s, schdlID)
	if err != nil {
		return logger.EchoHTTPError(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusNoContent)
}

func (g GameHandler) GetSchedule(c echo.Context) error {

	var gmeID int
	req := new(getScheduleReq)

	if err := req.bind(c, &gmeID); err != nil {
		return logger.EchoHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	currentSchdl, err := g.schdlStore.SelectScheduleByGameID(c.Request().Context(), gmeID, false)
	if err != nil {
		return logger.EchoHTTPError(http.StatusInternalServerError, err)
	}

	upcomingSchdl, err := g.schdlStore.SelectScheduleByGameID(c.Request().Context(), gmeID, true)
	if err != nil {
		return logger.EchoHTTPError(http.StatusInternalServerError, err)
	}

	var s model.ScheduleResp
	if currentSchdl != nil {
		s.WeekSchedule = currentSchdl.WkSchdl
	}
	if upcomingSchdl != nil {
		s.Upcoming = &model.UpcomingSchdl{WeekSchedule: upcomingSchdl.WkSchdl, ActiveBy: upcomingSchdl.ActiveBy}
	}

	return c.JSON(http.StatusOK, s)
}
