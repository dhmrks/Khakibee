package game_test

import (
	"encoding/json"
	"net/http"
	"strconv"
	"testing"

	"github.com/labstack/echo/v4"
	"gitlab.com/khakibee/khakibee/api/model"
	"gitlab.com/khakibee/khakibee/api/testUtils"
)

var mockShceduleReq = `{
  "0":["14:30","19:15","23:00"],
  "1":["14:30","19:15"],
  "2":["14:30","19:15"],
  "3":["14:30","19:15"],
  "4":["14:30","19:15"],
  "5":["14:30","19:15"],
  "6":["14:30","19:15","23:00"]  
}`

var mockNewShceduleReq = `{
  "0":["14:30","19:15","23:00"],
  "1":["14:30","19:15"],
  "2":["14:30","19:15"],
  "3":["14:30","19:15"],
  "4":["14:30","19:15"],
  "5":["14:30","19:15"],
  "6":["14:30","19:15","23:00"],
	"active_by":"2022-10-16"
}`

func TestPutNewSchedule(t *testing.T) {
	setup()

	gid := testUtils.DB1.Games[0].GameID
	gid2 := testUtils.DB1.Games[1].GameID

	schldID := new(int)
	*schldID = 2

	var schldIDNil *int

	t.Run("put schedule update successfully", func(t *testing.T) {
		c, _, res := testUtils.NewPutJSONRequest(mockNewShceduleReq)
		c.SetParamNames("gid")
		c.SetParamValues(strconv.Itoa(gid))

		err := h.PutNewSchedule(c)

		testUtils.AssertError(t, err, false)

		testUtils.AssertHTTPStatus(t, res.Code, http.StatusNoContent)
		testUtils.AssertFunctionCalleWith(t, ss.UpsertNewCalledWithSchdlID, schldID, "UpsertSchedule")
	})

	t.Run("put schedule insert successfully", func(t *testing.T) {
		c, _, res := testUtils.NewPutJSONRequest(mockNewShceduleReq)
		c.SetParamNames("gid")
		c.SetParamValues(strconv.Itoa(gid2))

		err := h.PutNewSchedule(c)

		testUtils.AssertError(t, err, false)

		testUtils.AssertHTTPStatus(t, res.Code, http.StatusNoContent)
		testUtils.AssertFunctionCalleWith(t, ss.UpsertNewCalledWithSchdlID, schldIDNil, "UpsertSchedule")
	})

	t.Run("put schedule with store error", func(t *testing.T) {
		c, _, _ := testUtils.NewPutJSONRequest(mockNewShceduleReq)
		c.SetParamNames("gid")
		c.SetParamValues(strconv.Itoa(gid))

		ss.ReturnError = errStore

		err := h.PutNewSchedule(c).(*echo.HTTPError)

		testUtils.AssertError(t, err, true)

		testUtils.AssertHTTPStatus(t, err.Code, http.StatusInternalServerError)
	})
}

func TestPutSchedule(t *testing.T) {
	setup()

	gid := testUtils.DB1.Games[0].GameID
	gid2 := testUtils.DB1.Games[2].GameID

	schldID := new(int)
	*schldID = 1

	var schldIDNil *int

	t.Run("put schedule update successfully", func(t *testing.T) {
		c, _, res := testUtils.NewPutJSONRequest(mockShceduleReq)
		c.SetParamNames("gid")
		c.SetParamValues(strconv.Itoa(gid))

		err := h.PutSchedule(c)

		testUtils.AssertError(t, err, false)

		testUtils.AssertHTTPStatus(t, res.Code, http.StatusNoContent)
		testUtils.AssertFunctionCalleWith(t, ss.UpsertCalledWithSchdlID, schldID, "UpsertSchedule")
	})

	t.Run("put schedule insert successfully", func(t *testing.T) {
		c, _, res := testUtils.NewPutJSONRequest(mockShceduleReq)
		c.SetParamNames("gid")
		c.SetParamValues(strconv.Itoa(gid2))

		err := h.PutSchedule(c)

		testUtils.AssertError(t, err, false)

		testUtils.AssertHTTPStatus(t, res.Code, http.StatusNoContent)
		testUtils.AssertFunctionCalleWith(t, ss.UpsertCalledWithSchdlID, schldIDNil, "UpsertSchedule")
	})

	t.Run("put schedule with store error", func(t *testing.T) {
		c, _, _ := testUtils.NewPutJSONRequest(mockShceduleReq)
		c.SetParamNames("gid")
		c.SetParamValues(strconv.Itoa(gid))

		ss.ReturnError = errStore

		err := h.PutSchedule(c).(*echo.HTTPError)

		testUtils.AssertError(t, err, true)

		testUtils.AssertHTTPStatus(t, err.Code, http.StatusInternalServerError)
	})
}

func TestGetSchedule(t *testing.T) {
	setup()

	gid := testUtils.DB1.Games[0].GameID
	gid2 := testUtils.DB1.Games[1].GameID

	wantSchdlWithUpc := model.ScheduleResp{
		WeekSchedule: testUtils.DB1.Schedule[gid][0].WkSchdl,
		Upcoming: &model.UpcomingSchdl{
			WeekSchedule: testUtils.DB1.Schedule[gid][1].WkSchdl,
			ActiveBy:     testUtils.DB1.Schedule[gid][1].ActiveBy,
		},
	}

	wantSchdlWithoutUpc := model.ScheduleResp{
		WeekSchedule: testUtils.DB1.Schedule[gid2][0].WkSchdl,
	}

	t.Run("get schedule with upcoming successfully", func(t *testing.T) {
		c, _, res := testUtils.NewGetRequest()
		c.SetParamNames("gid")
		c.SetParamValues(strconv.Itoa(gid))

		err := h.GetSchedule(c)

		var gotSchdl model.ScheduleResp
		json.NewDecoder(res.Body).Decode(&gotSchdl)

		testUtils.AssertError(t, err, false)
		testUtils.AssertHTTPStatus(t, res.Code, http.StatusOK)
		testUtils.AssertFunctionCalled(t, ss.SelectCalled, 2, "SelectScheduleByGameID")

		testUtils.AssertEqual(t, gotSchdl, wantSchdlWithUpc)

	})

	t.Run("get schedule without upcoming successfully", func(t *testing.T) {
		c, _, res := testUtils.NewGetRequest()
		c.SetParamNames("gid")
		c.SetParamValues(strconv.Itoa(gid2))

		err := h.GetSchedule(c)

		var gotSchdl model.ScheduleResp
		json.NewDecoder(res.Body).Decode(&gotSchdl)

		testUtils.AssertError(t, err, false)
		testUtils.AssertHTTPStatus(t, res.Code, http.StatusOK)
		testUtils.AssertFunctionCalled(t, ss.SelectCalled, 4, "SelectScheduleByGameID")
		testUtils.AssertEqual(t, gotSchdl, wantSchdlWithoutUpc)

	})

	t.Run("get schedule empty schedule", func(t *testing.T) {
		c, _, res := testUtils.NewGetRequest()
		c.SetParamNames("gid")
		c.SetParamValues(strconv.Itoa(3))

		err := h.GetSchedule(c)

		var gotSchdl model.ScheduleResp
		json.NewDecoder(res.Body).Decode(&gotSchdl)

		testUtils.AssertError(t, err, false)
		testUtils.AssertHTTPStatus(t, res.Code, http.StatusOK)
		testUtils.AssertFunctionCalled(t, ss.SelectCalled, 6, "SelectScheduleByGameID")
		testUtils.AssertEqual(t, gotSchdl, model.ScheduleResp{})
	})

	t.Run("get schedule with store error", func(t *testing.T) {
		c, _, _ := testUtils.NewGetRequest()
		c.SetParamNames("gid")
		c.SetParamValues(strconv.Itoa(3))

		ss.ReturnError = errStore

		err := h.GetSchedule(c).(*echo.HTTPError)

		testUtils.AssertError(t, err, true)
		testUtils.AssertHTTPStatus(t, err.Code, http.StatusInternalServerError)
		testUtils.AssertFunctionCalled(t, ss.SelectCalled, 7, "SelectScheduleByGameID")
	})

}
