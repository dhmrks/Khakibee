package game_test

import (
	"encoding/json"
	"net/http"
	"strconv"
	"testing"

	"github.com/labstack/echo/v4"
	"gitlab.com/khakibee/khakibee/api/email"
	"gitlab.com/khakibee/khakibee/api/model"
	"gitlab.com/khakibee/khakibee/api/testUtils"
	"gitlab.com/khakibee/khakibee/api/user"
)

var (
	mockBookingReqs = []string{
		`{
			"status":"booked",
			"date":"2022-06-30 14:00",
			"name":"ioulios",
			"mob_num":"+306975645865",
			"email":"ioulis@email.com"
		}`,
		`{
			"status":"booked",
			"date":"2022-06-29 19:00",
			"name":"ioulios",
			"mob_num":"+306975645865",
			"email":"ioulis@email.com"
		}`,
		//out of schedule case
		`{
			"status":"booked",
			"date":"2022-06-30 16:00",
			"name":"ioulios",
			"mob_num":"+306975645865",
			"email":"ioulis@email.com"
		}`,
	}
)

func TestRemoveBooking(t *testing.T) {
	setup()
	gid := strconv.Itoa(testUtils.DB1.Bookings[0].GameID)
	bid := strconv.Itoa(testUtils.DB1.Bookings[0].BkngID)

	wantEmailParams := email.SendEmailParams{BkngID: testUtils.DB1.Bookings[0].BkngID, TmpltCode: email.TempleteCancelBkng, Timezone: testUtils.Timezone}

	t.Run("cancel booking successfully", func(t *testing.T) {
		c, _, res := testUtils.NewDeleteRequest()
		c.SetParamNames("gid", "bid")
		c.SetParamValues(gid, bid)

		h.RemoveBooking(c)

		testUtils.AssertFunctionCalled(t, bs.CancelCalled, 1, "CancelBooking")
		testUtils.AssertFunctionCalleWith(t, eh.SendEmailCalledWith, wantEmailParams, "SendEmail")
		testUtils.AssertHTTPStatus(t, res.Code, http.StatusNoContent)
	})

	t.Run("cancel booking successfully with email error", func(t *testing.T) {
		bs.CancelCalled = 0
		c, _, res := testUtils.NewDeleteRequest()
		c.SetParamNames("gid", "bid")
		c.SetParamValues(gid, bid)

		eh.ReturnError = errStore

		h.RemoveBooking(c)

		testUtils.AssertFunctionCalled(t, bs.CancelCalled, 1, "CancelBooking")
		testUtils.AssertFunctionCalleWith(t, eh.SendEmailCalledWith, wantEmailParams, "SendEmail")
		testUtils.AssertHTTPStatus(t, res.Code, http.StatusNoContent)
	})

	t.Run("cancel booking not found", func(t *testing.T) {
		bs.CancelCalled = 0

		c, _, _ := testUtils.NewDeleteRequest()
		c.SetParamNames("gid", "bid")
		c.SetParamValues(gid, "0")

		err := h.RemoveBooking(c).(*echo.HTTPError)

		testUtils.AssertError(t, err, true)
		testUtils.AssertFunctionCalled(t, bs.CancelCalled, 1, "CancelBooking")
		testUtils.AssertHTTPStatus(t, err.Code, http.StatusNotFound)
	})

	t.Run("cancel booking with params", func(t *testing.T) {
		bs.CancelCalled = 0
		c, _, _ := testUtils.NewDeleteRequest()
		c.SetParamNames("gid", "bid")
		c.SetParamValues(gid, "a")

		err := h.RemoveBooking(c).(*echo.HTTPError)

		testUtils.AssertFunctionCalled(t, bs.CancelCalled, 0, "CancelBooking")
		testUtils.AssertHTTPStatus(t, err.Code, http.StatusUnprocessableEntity)
	})

	t.Run("cancel booking with failing call in store", func(t *testing.T) {
		bs.CancelCalled = 0
		c, _, _ := testUtils.NewDeleteRequest()
		c.SetParamNames("gid", "bid")
		c.SetParamValues(gid, bid)

		bs.ReturnError = errStore

		err := h.RemoveBooking(c).(*echo.HTTPError)

		testUtils.AssertFunctionCalled(t, bs.CancelCalled, 1, "CancelBooking")
		testUtils.AssertHTTPStatus(t, err.Code, http.StatusInternalServerError)
		testUtils.AssertEqual(t, err.Message, errStore.Error())
	})
}

func TestEditBooking(t *testing.T) {
	setup()
	gid := strconv.Itoa(testUtils.DB1.Bookings[0].GameID)
	bid := strconv.Itoa(testUtils.DB1.Bookings[0].BkngID)
	bid2 := strconv.Itoa(testUtils.DB1.Bookings[1].BkngID)
	uid := testUtils.DB1.Users[0].UserID

	wantEmailParams := email.SendEmailParams{BkngID: testUtils.DB1.Bookings[0].BkngID, TmpltCode: email.TempleteUpdateBkng, Timezone: testUtils.Timezone}

	t.Run("edit booking successfully", func(t *testing.T) {
		c, _, res := testUtils.NewPutJSONRequest(mockBookingReqs[0])
		c.SetParamNames("gid", "bid")
		c.SetParamValues(gid, bid)

		c.Set(user.UserIDCtxKey, uid)

		err := h.EditBooking(c)

		testUtils.AssertError(t, err, false)
		testUtils.AssertHTTPStatus(t, res.Code, http.StatusNoContent)
		testUtils.AssertFunctionCalled(t, ss.CheckCalled, 1, "CheckSchedule")
		testUtils.AssertFunctionCalled(t, bs.UpdateCalled, 1, "UpdateBooking")
		testUtils.AssertFunctionCalleWith(t, bs.UpdateCalledWithID, uid, "UpdateBooking")

		testUtils.AssertFunctionCalleWith(t, eh.SendEmailCalledWith, wantEmailParams, "SendEmail")
	})

	t.Run("edit booking successfully with email error", func(t *testing.T) {
		ss.CheckCalled = 0
		bs.UpdateCalled = 0
		eh.ReturnError = errStore

		c, _, res := testUtils.NewPutJSONRequest(mockBookingReqs[0])
		c.SetParamNames("gid", "bid")
		c.SetParamValues(gid, bid)

		c.Set(user.UserIDCtxKey, uid)

		err := h.EditBooking(c)

		testUtils.AssertError(t, err, false)
		testUtils.AssertHTTPStatus(t, res.Code, http.StatusNoContent)
		testUtils.AssertFunctionCalled(t, ss.CheckCalled, 1, "CheckSchedule")
		testUtils.AssertFunctionCalled(t, bs.UpdateCalled, 1, "UpdateBooking")
		testUtils.AssertFunctionCalleWith(t, bs.UpdateCalledWithID, uid, "UpdateBooking")

		testUtils.AssertFunctionCalleWith(t, eh.SendEmailCalledWith, wantEmailParams, "SendEmail")
	})

	t.Run("edit booking with invalid json", func(t *testing.T) {
		ss.CheckCalled = 0
		bs.UpdateCalled = 0

		c, _, _ := testUtils.NewPutJSONRequest(``)
		c.SetParamNames("gid", "bid")
		c.SetParamValues(gid, bid)

		c.Set(user.UserIDCtxKey, uid)

		err := h.EditBooking(c).(*echo.HTTPError)
		testUtils.AssertError(t, err, true)
		testUtils.AssertHTTPStatus(t, err.Code, http.StatusUnprocessableEntity)

		testUtils.AssertFunctionCalled(t, ss.CheckCalled, 0, "CheckSchedule")
		testUtils.AssertFunctionCalled(t, bs.UpdateCalled, 0, "UpdateBooking")
	})

	t.Run("edit booking on non existent date", func(t *testing.T) {
		ss.CheckCalled = 0
		bs.UpdateCalled = 0

		c, _, _ := testUtils.NewPutJSONRequest(mockBookingReqs[2])
		c.SetParamNames("gid", "bid")
		c.SetParamValues(gid, bid)

		c.Set(user.UserIDCtxKey, uid)

		err := h.EditBooking(c).(*echo.HTTPError)
		testUtils.AssertError(t, err, true)
		testUtils.AssertHTTPStatus(t, err.Code, http.StatusBadRequest)

		testUtils.AssertFunctionCalled(t, ss.CheckCalled, 1, "CheckSchedule")
		testUtils.AssertFunctionCalled(t, bs.UpdateCalled, 0, "UpdateBooking")
	})

	t.Run("edit booking with unavailable date", func(t *testing.T) {
		ss.CheckCalled = 0
		bs.UpdateCalled = 0

		c, _, _ := testUtils.NewPutJSONRequest(mockBookingReqs[0])
		c.SetParamNames("gid", "bid")
		c.SetParamValues(gid, bid2)

		c.Set(user.UserIDCtxKey, uid)

		err := h.EditBooking(c).(*echo.HTTPError)
		testUtils.AssertError(t, err, true)
		testUtils.AssertHTTPStatus(t, err.Code, http.StatusConflict)
		testUtils.AssertFunctionCalled(t, ss.CheckCalled, 1, "CheckSchedule")
		testUtils.AssertFunctionCalled(t, bs.UpdateCalled, 1, "UpdateBooking")
	})

	t.Run("edit booking not found", func(t *testing.T) {
		ss.CheckCalled = 0
		bs.UpdateCalled = 0

		c, _, _ := testUtils.NewPutJSONRequest(mockBookingReqs[0])
		c.SetParamNames("gid", "bid")
		c.SetParamValues(gid, "0")

		c.Set(user.UserIDCtxKey, uid)

		err := h.EditBooking(c).(*echo.HTTPError)
		testUtils.AssertError(t, err, true)
		testUtils.AssertHTTPStatus(t, err.Code, http.StatusNotFound)
		testUtils.AssertFunctionCalled(t, ss.CheckCalled, 1, "CheckSchedule")
		testUtils.AssertFunctionCalled(t, bs.UpdateCalled, 1, "UpdateBooking")
	})

	t.Run("edit booking with error game store", func(t *testing.T) {
		ss.CheckCalled = 0
		bs.UpdateCalled = 0
		bs.ReturnError = errStore

		c, _, _ := testUtils.NewPutJSONRequest(mockBookingReqs[0])
		c.SetParamNames("gid", "bid")
		c.SetParamValues(gid, bid)

		c.Set(user.UserIDCtxKey, uid)

		err := h.EditBooking(c).(*echo.HTTPError)
		testUtils.AssertError(t, err, true)
		testUtils.AssertHTTPStatus(t, err.Code, http.StatusInternalServerError)
		testUtils.AssertFunctionCalled(t, ss.CheckCalled, 1, "CheckSchedule")
		testUtils.AssertFunctionCalled(t, bs.UpdateCalled, 1, "UpdateBooking")
	})

	t.Run("edit booking with error booking store", func(t *testing.T) {
		ss.CheckCalled = 0
		bs.UpdateCalled = 0
		gs.ReturnError = nil
		bs.ReturnError = errStore

		c, _, _ := testUtils.NewPutJSONRequest(mockBookingReqs[0])
		c.SetParamNames("gid", "bid")
		c.SetParamValues(gid, bid)

		c.Set(user.UserIDCtxKey, uid)

		err := h.EditBooking(c).(*echo.HTTPError)
		testUtils.AssertError(t, err, true)
		testUtils.AssertHTTPStatus(t, err.Code, http.StatusInternalServerError)
		testUtils.AssertFunctionCalled(t, ss.CheckCalled, 1, "CheckSchedule")
		testUtils.AssertFunctionCalled(t, bs.UpdateCalled, 1, "UpdateBooking")
	})

}

func TestGetBooking(t *testing.T) {
	setup()
	gid := strconv.Itoa(testUtils.DB1.Bookings[0].GameID)
	bid := strconv.Itoa(testUtils.DB1.Bookings[0].BkngID)

	t.Run("get booking successfully", func(t *testing.T) {
		c, _, res := testUtils.NewGetRequest()
		c.SetParamNames("gid", "bid")
		c.SetParamValues(gid, bid)

		h.GetBooking(c)
		testUtils.AssertHTTPStatus(t, res.Code, http.StatusOK)

		var gotBkng model.Booking
		json.NewDecoder(res.Body).Decode(&gotBkng)

		testUtils.AssertEqual(t, gotBkng, testUtils.DB1.Bookings[0])
	})

	t.Run("get booking not found", func(t *testing.T) {
		c, _, _ := testUtils.NewGetRequest()
		c.SetParamNames("gid", "bid")
		c.SetParamValues(gid, "0")

		err := h.GetBooking(c).(*echo.HTTPError)
		testUtils.AssertHTTPStatus(t, err.Code, http.StatusNotFound)
		testUtils.AssertError(t, err, true)
	})

	t.Run("get booking with sql error", func(t *testing.T) {
		bs.ReturnError = errStore

		c, _, _ := testUtils.NewGetRequest()
		c.SetParamNames("gid", "bid")
		c.SetParamValues(gid, bid)

		err := h.GetBooking(c).(*echo.HTTPError)

		testUtils.AssertHTTPStatus(t, err.Code, http.StatusInternalServerError)
		testUtils.AssertError(t, err, true)
	})

	t.Run("get booking invalid id type", func(t *testing.T) {
		bs.ReturnError = nil
		c, _, _ := testUtils.NewGetRequest()
		c.SetParamNames("gid", "bid")
		c.SetParamValues(gid, "a")

		err := h.GetBooking(c).(*echo.HTTPError)

		testUtils.AssertHTTPStatus(t, err.Code, http.StatusUnprocessableEntity)
		testUtils.AssertError(t, err, true)
	})
}

func TestGetBookings(t *testing.T) {
	setup()
	gid := strconv.Itoa(testUtils.DB1.Bookings[0].GameID)
	var gid1bs []model.Booking
	for idx, b := range testUtils.DB1.Bookings {
		if strconv.Itoa(b.GameID) == gid {
			gid1bs = append(gid1bs, testUtils.DB1.Bookings[idx])
		}
	}

	t.Run("get bookings successfully", func(t *testing.T) {
		c, _, res := testUtils.NewGetRequest()

		c.SetParamNames("gid")
		c.SetParamValues(gid)

		h.GetBookings(c)
		testUtils.AssertHTTPStatus(t, res.Code, http.StatusOK)

		var gotBkngs []model.Booking
		json.NewDecoder(res.Body).Decode(&gotBkngs)

		testUtils.AssertEqual(t, gotBkngs, gid1bs)
	})

	t.Run("get bookings empty", func(t *testing.T) {
		c, _, res := testUtils.NewGetRequest()
		c.SetParamNames("gid")
		c.SetParamValues("0")

		h.GetBookings(c)
		testUtils.AssertHTTPStatus(t, res.Code, http.StatusOK)

		var gotBkngs []model.Booking
		json.NewDecoder(res.Body).Decode(&gotBkngs)

		testUtils.AssertEqual(t, len(gotBkngs), 0)
	})

	t.Run("get bookings with sql error", func(t *testing.T) {
		bs.ReturnError = errStore
		c, _, _ := testUtils.NewGetRequest()
		c.SetParamNames("gid")
		c.SetParamValues(gid)

		err := h.GetBookings(c).(*echo.HTTPError)
		testUtils.AssertHTTPStatus(t, err.Code, http.StatusInternalServerError)
		testUtils.AssertError(t, err, true)
	})
}

func TestBookGame(t *testing.T) {
	setup()
	gid := strconv.Itoa(testUtils.DB1.Bookings[0].GameID)
	uid := testUtils.DB1.Users[0].UserID

	wantEmailParams := email.SendEmailParams{BkngID: testUtils.DB1.Bookings[0].BkngID, TmpltCode: email.TempleteNewBkng, Timezone: testUtils.Timezone}

	t.Run("book game successfully", func(t *testing.T) {

		bs.NewBookingID = testUtils.DB1.Bookings[0].BkngID

		c, _, res := testUtils.NewPostJSONRequest(mockBookingReqs[1])
		c.SetParamNames("gid")
		c.SetParamValues(gid)

		c.Set(user.UserIDCtxKey, uid)

		err := h.BookGame(c)
		testUtils.AssertError(t, err, false)
		testUtils.AssertHTTPStatus(t, res.Code, http.StatusCreated)
		testUtils.AssertFunctionCalled(t, ss.CheckCalled, 1, "CheckSchedule")
		testUtils.AssertFunctionCalled(t, bs.InsertCalled, 1, "InsertBooking")
		testUtils.AssertFunctionCalleWith(t, bs.InsertCalledWithID, uid, "InsertBooking")

		testUtils.AssertFunctionCalleWith(t, eh.SendEmailCalledWith, wantEmailParams, "SendEmail")
	})

	t.Run("book game successfully with email error", func(t *testing.T) {

		ss.CheckCalled = 0
		bs.InsertCalled = 0
		bs.NewBookingID = testUtils.DB1.Bookings[0].BkngID
		eh.ReturnError = errStore

		c, _, res := testUtils.NewPostJSONRequest(mockBookingReqs[1])
		c.SetParamNames("gid")
		c.SetParamValues(gid)

		c.Set(user.UserIDCtxKey, uid)

		err := h.BookGame(c)
		testUtils.AssertError(t, err, false)
		testUtils.AssertHTTPStatus(t, res.Code, http.StatusCreated)
		testUtils.AssertFunctionCalled(t, ss.CheckCalled, 1, "CheckSchedule")
		testUtils.AssertFunctionCalled(t, bs.InsertCalled, 1, "InsertBooking")
		testUtils.AssertFunctionCalleWith(t, bs.InsertCalledWithID, uid, "InsertBooking")

		testUtils.AssertFunctionCalleWith(t, eh.SendEmailCalledWith, wantEmailParams, "SendEmail")
	})

	t.Run("Book game with invalid json", func(t *testing.T) {
		ss.CheckCalled = 0
		bs.InsertCalled = 0

		c, _, _ := testUtils.NewPostJSONRequest(``)

		c.Set(user.UserIDCtxKey, uid)

		c.SetParamNames("gid")
		c.SetParamValues(gid)

		err := h.BookGame(c).(*echo.HTTPError)
		testUtils.AssertError(t, err, true)
		testUtils.AssertHTTPStatus(t, err.Code, http.StatusUnprocessableEntity)

		testUtils.AssertFunctionCalled(t, ss.CheckCalled, 0, "CheckSchedule")
		testUtils.AssertFunctionCalled(t, bs.InsertCalled, 0, "InsertBooking")
		testUtils.AssertFunctionCalleWith(t, bs.InsertCalledWithID, uid, "InsertBooking")
	})

	t.Run("Book game out of schechule", func(t *testing.T) {
		ss.CheckCalled = 0
		bs.InsertCalled = 0

		c, _, _ := testUtils.NewPostJSONRequest(mockBookingReqs[2])
		c.SetParamNames("gid")
		c.SetParamValues(gid)

		c.Set(user.UserIDCtxKey, uid)

		err := h.BookGame(c).(*echo.HTTPError)
		testUtils.AssertError(t, err, true)
		testUtils.AssertHTTPStatus(t, err.Code, http.StatusNotFound)

		testUtils.AssertFunctionCalled(t, ss.CheckCalled, 1, "CheckSchedule")
		testUtils.AssertFunctionCalled(t, bs.InsertCalled, 0, "InsertBooking")
		testUtils.AssertFunctionCalleWith(t, bs.InsertCalledWithID, uid, "InsertBooking")
	})

	t.Run("Book game on unavailable date", func(t *testing.T) {
		ss.CheckCalled = 0
		bs.InsertCalled = 0

		c, _, _ := testUtils.NewPostJSONRequest(mockBookingReqs[0])
		c.SetParamNames("gid")
		c.SetParamValues(gid)

		c.Set(user.UserIDCtxKey, uid)

		err := h.BookGame(c).(*echo.HTTPError)
		testUtils.AssertError(t, err, true)
		testUtils.AssertHTTPStatus(t, err.Code, http.StatusConflict)
		testUtils.AssertFunctionCalled(t, ss.CheckCalled, 1, "CheckSchedule")
		testUtils.AssertFunctionCalled(t, bs.InsertCalled, 1, "InsertBooking")
		testUtils.AssertFunctionCalleWith(t, bs.InsertCalledWithID, uid, "InsertBooking")
	})

	t.Run("Book game catch game store error", func(t *testing.T) {
		ss.CheckCalled = 0
		bs.InsertCalled = 0
		bs.ReturnError = errStore

		c, _, _ := testUtils.NewPostJSONRequest(mockBookingReqs[0])
		c.SetParamNames("gid")
		c.SetParamValues(gid)

		c.Set(user.UserIDCtxKey, uid)

		err := h.BookGame(c).(*echo.HTTPError)
		testUtils.AssertError(t, err, true)
		testUtils.AssertHTTPStatus(t, err.Code, http.StatusInternalServerError)
		testUtils.AssertFunctionCalled(t, ss.CheckCalled, 1, "CheckSchedule")
		testUtils.AssertFunctionCalled(t, bs.InsertCalled, 1, "InsertBooking")
		testUtils.AssertFunctionCalleWith(t, bs.InsertCalledWithID, uid, "InsertBooking")
	})

	t.Run("Book game catch booking store error", func(t *testing.T) {
		ss.CheckCalled = 0
		bs.InsertCalled = 0
		gs.ReturnError = nil
		bs.ReturnError = errStore

		c, _, _ := testUtils.NewPostJSONRequest(mockBookingReqs[0])
		c.SetParamNames("gid")
		c.SetParamValues(gid)

		c.Set(user.UserIDCtxKey, uid)

		err := h.BookGame(c).(*echo.HTTPError)
		testUtils.AssertError(t, err, true)
		testUtils.AssertHTTPStatus(t, err.Code, http.StatusInternalServerError)
		testUtils.AssertFunctionCalled(t, ss.CheckCalled, 1, "CheckSchedule")
		testUtils.AssertFunctionCalled(t, bs.InsertCalled, 1, "InsertBooking")
		testUtils.AssertFunctionCalleWith(t, bs.InsertCalledWithID, uid, "InsertBooking")
	})

}
