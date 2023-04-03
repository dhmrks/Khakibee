package game

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"gitlab.com/khakibee/khakibee/api/email"
	"gitlab.com/khakibee/khakibee/api/logger"
	"gitlab.com/khakibee/khakibee/api/model"
	"gitlab.com/khakibee/khakibee/api/user"
)

func (b *GameHandler) RemoveBooking(c echo.Context) error {
	var (
		bk  model.Booking
		loc time.Location
	)
	req := new(removeBookingReq)

	if err := req.bind(c, &bk, &loc); err != nil {
		return logger.EchoHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	fmt.Println("loc", loc.String())

	ok, err := b.bkngStore.CancelBooking(c.Request().Context(), bk.GameID, bk.BkngID)
	if err != nil {
		return logger.EchoHTTPError(http.StatusInternalServerError, err)
	}
	if !ok {
		return logger.EchoHTTPError(http.StatusNotFound, "Booking not found")
	}

	c.Response().WriteHeader(http.StatusNoContent)
	c.Response().Flush()

	if err = b.emailHandler.SendEmail(bk.BkngID, email.TempleteCancelBkng, loc.String()); err != nil {
		c.Logger().Error(err)
	}

	return nil

}

func (b *GameHandler) EditBooking(c echo.Context) error {
	var (
		bk  model.Booking
		loc time.Location
	)
	req := new(editBookingReq)

	if err := req.bind(c, &bk, &loc); err != nil {
		return logger.EchoHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	exists, err := b.schdlStore.CheckSchedule(c.Request().Context(), bk.Dte, bk.GameID)
	if err != nil {
		return logger.EchoHTTPError(http.StatusInternalServerError, err.Error())
	}
	if !exists {
		return logger.EchoHTTPError(http.StatusBadRequest, "Booking date for selected game does not exists ")
	}

	userID := c.Get(user.UserIDCtxKey).(int)

	notFound, dateConflict, err := b.bkngStore.UpdateBooking(c.Request().Context(), bk, userID)
	if err != nil {
		return logger.EchoHTTPError(http.StatusInternalServerError, err.Error())
	}
	if notFound {
		return logger.EchoHTTPError(http.StatusNotFound, "Booking not found")
	}
	if dateConflict {
		return logger.EchoHTTPError(http.StatusConflict, "Selected date is unavailable")
	}

	c.Response().WriteHeader(http.StatusNoContent)
	c.Response().Flush()

	if err = b.emailHandler.SendEmail(bk.BkngID, email.TempleteUpdateBkng, loc.String()); err != nil {
		c.Logger().Error(err)
	}

	return nil
}

func (b *GameHandler) GetBooking(c echo.Context) error {
	var bk model.Booking
	req := new(getBookingReq)

	if err := req.bind(c, &bk); err != nil {
		return logger.EchoHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	bkng, err := b.bkngStore.SelectBookingByID(c.Request().Context(), bk.GameID, bk.BkngID)
	if err != nil {
		return logger.EchoHTTPError(http.StatusInternalServerError, err)
	} else if bkng == nil {
		return logger.EchoHTTPError(http.StatusNotFound, "Booking not found")
	}
	return c.JSON(http.StatusOK, bkng)
}

func (b *GameHandler) GetBookings(c echo.Context) error {
	var (
		bk  model.Booking
		loc time.Location
	)
	req := new(getBookingsReq)

	if err := req.bind(c, &bk, &loc); err != nil {
		return logger.EchoHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	bkngs, err := b.bkngStore.SelectBookings(c.Request().Context(), bk.GameID, loc.String(), 0)
	if err != nil {
		return logger.EchoHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, bkngs)
}

func (b *GameHandler) BookGame(c echo.Context) error {
	var (
		bk  model.Booking
		loc time.Location
	)
	req := new(bookGameReq)

	if err := req.bind(c, &bk, &loc); err != nil {
		return logger.EchoHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	userID := c.Get(user.UserIDCtxKey).(int)

	exists, err := b.schdlStore.CheckSchedule(c.Request().Context(), req.Dte, bk.GameID)
	if err != nil {
		return logger.EchoHTTPError(http.StatusInternalServerError, err.Error())
	}
	if !exists {
		return logger.EchoHTTPError(http.StatusNotFound, "Booking date for selected game does not exists ")
	}

	newBkngID, err := b.bkngStore.InsertBooking(c.Request().Context(), bk, userID)
	if err != nil {
		return logger.EchoHTTPError(http.StatusInternalServerError, err.Error())
	}
	if newBkngID == nil {
		return logger.EchoHTTPError(http.StatusConflict, "Selected date is unavailable")
	}

	c.Response().WriteHeader(http.StatusCreated)
	c.Response().Flush()

	if err = b.emailHandler.SendEmail(*newBkngID, email.TempleteNewBkng, loc.String()); err != nil {
		c.Logger().Error(err)
	}

	return nil

}
