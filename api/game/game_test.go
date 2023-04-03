package game_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"github.com/labstack/echo/v4"

	"gitlab.com/khakibee/khakibee/api/email"
	"gitlab.com/khakibee/khakibee/api/game"
	"gitlab.com/khakibee/khakibee/api/model"
	"gitlab.com/khakibee/khakibee/api/store"
	"gitlab.com/khakibee/khakibee/api/testUtils"
	"gitlab.com/khakibee/khakibee/api/user"
)

var (
	gs       *store.MockGameStore
	bs       *store.MockBookingStore
	ss       *store.MockScheduleStore
	eh       *email.MockEmailHandler
	h        *game.GameHandler
	errStore = errors.New("internal server error")
)

func setup() {
	gs = &store.MockGameStore{Games: testUtils.DB1.Games, Calendar: testUtils.DB1.Calendar}
	bs = &store.MockBookingStore{Bookings: testUtils.DB1.Bookings}
	ss = &store.MockScheduleStore{Calendar: testUtils.DB1.Calendar, Schedule: testUtils.DB1.Schedule}

	eh = email.NewMockEmailHandler()

	h = game.NewGameHandler(gs, bs, ss, eh)
}

var (
	mockGameReqs = []string{`
		{
			"status": "active",
			"name":     "Music Academy",
			"descr":    "A music academy",
			"addr":     "21 Athens str, Athens, Greece 21012",
			"img_url":  "https://booking.com/musicacademy.jpg",
			"map_url":  "https://goo.gl/maps/ghrrS37V7tPGw5Zd8",
			"players":  "2-5",
			"duration": 150,
			"age_range":"13+"
		}`,
		`{
			"status": "active",
			"name":     "Music's Academy",
			"descr":    "A music academy !@#$%^&*()_=-:",
			"addr":     "21 Athens str, Athens, Greece 21012",
			"img_url":  "https://booking.com",
			"map_url":  "https://www.goo.gl/maps/ghrrS37V7tPGw5Zd8",
			"players":  "2-5",
			"duration":  150,
			"age_range":"13-18"
		}`,
	}
)

func TestEditGame(t *testing.T) {
	setup()

	t.Run("edit game successfully", func(t *testing.T) {
		c, _, res := testUtils.NewPutJSONRequest(mockGameReqs[0])
		c.SetParamNames("gid")
		c.SetParamValues(strconv.Itoa(testUtils.DB1.Games[0].GameID))

		err := h.EditGame(c)

		testUtils.AssertError(t, err, false)

		testUtils.AssertHTTPStatus(t, res.Code, http.StatusNoContent)
		testUtils.AssertFunctionCalled(t, gs.UpdateCalled, 1, "UpdateGame")
	})

	t.Run("edit game not found", func(t *testing.T) {
		gs.UpdateCalled = 0
		c, _, _ := testUtils.NewPutJSONRequest(mockGameReqs[0])
		c.SetParamNames("gid")
		c.SetParamValues("0")

		err := h.EditGame(c)
		he := err.(*echo.HTTPError)

		testUtils.AssertHTTPStatus(t, he.Code, http.StatusNotFound)
		testUtils.AssertFunctionCalled(t, gs.UpdateCalled, 1, "UpdateGame")
	})

	t.Run("edit game with invalid json", func(t *testing.T) {
		c, _, _ := testUtils.NewPostJSONRequest("{}")
		c.SetParamNames("gid")
		c.SetParamValues(strconv.Itoa(testUtils.DB1.Games[0].GameID))

		err := h.EditGame(c)
		he := err.(*echo.HTTPError)

		testUtils.AssertHTTPStatus(t, he.Code, http.StatusUnprocessableEntity)
	})

	t.Run("edit game with failing update in store", func(t *testing.T) {
		c, _, _ := testUtils.NewPostJSONRequest(mockGameReqs[0])
		c.SetParamNames("gid")
		c.SetParamValues(strconv.Itoa(testUtils.DB1.Games[0].GameID))

		gs.ReturnError = errStore

		err := h.EditGame(c).(*echo.HTTPError)

		testUtils.AssertHTTPStatus(t, err.Code, http.StatusInternalServerError)
		testUtils.AssertEqual(t, err.Message, gs.ReturnError.Error())
	})

}

func TestGetGame(t *testing.T) {
	setup()

	for _, tg := range testUtils.DB1.Games {
		t.Run(fmt.Sprintf("Get game %d successfully", tg.GameID), func(t *testing.T) {
			c, _, res := testUtils.NewGetRequest()
			c.SetParamNames("gid")
			c.SetParamValues(strconv.Itoa(tg.GameID))

			h.GetGame(c)
			testUtils.AssertHTTPStatus(t, res.Code, http.StatusOK)

			var gotGame model.Game
			json.NewDecoder(res.Body).Decode(&gotGame)

			testUtils.AssertEqual(t, gotGame, tg)

		})
	}

	t.Run("Get game not found", func(t *testing.T) {
		c, _, _ := testUtils.NewGetRequest()
		c.SetParamNames("gid")
		c.SetParamValues("0")

		err := h.GetGame(c)
		he := err.(*echo.HTTPError)

		testUtils.AssertHTTPStatus(t, he.Code, http.StatusNotFound)
	})

	t.Run("Get game with invalid id type", func(t *testing.T) {
		c, _, _ := testUtils.NewGetRequest()
		c.SetParamNames("gid")
		c.SetParamValues("a")

		err := h.GetGame(c)
		he := err.(*echo.HTTPError)

		testUtils.AssertHTTPStatus(t, he.Code, http.StatusUnprocessableEntity)
	})

	t.Run("Get game with failing select in store", func(t *testing.T) {
		c, _, _ := testUtils.NewGetRequest()
		c.SetParamNames("gid")
		c.SetParamValues(strconv.Itoa(testUtils.DB1.Games[0].GameID))

		gs.ReturnError = errStore

		err := h.GetGame(c)
		he := err.(*echo.HTTPError)

		testUtils.AssertHTTPStatus(t, he.Code, http.StatusInternalServerError)
	})

}

func TestCreateGame(t *testing.T) {
	setup()
	for idx, mng := range mockGameReqs {
		t.Run(fmt.Sprintf("Create new game successfully %d", idx), func(t *testing.T) {
			gs.InsertCalled = 0
			c, _, res := testUtils.NewPostJSONRequest(mng)

			h.CreateGame(c)

			testUtils.AssertHTTPStatus(t, res.Code, http.StatusCreated)
			testUtils.AssertFunctionCalled(t, gs.InsertCalled, 1, "InsertGame")
		})
	}

	t.Run("Create new game with invalid json values", func(t *testing.T) {
		c, _, _ := testUtils.NewPostJSONRequest("{}")

		err := h.CreateGame(c)
		he := err.(*echo.HTTPError)

		testUtils.AssertHTTPStatus(t, he.Code, http.StatusUnprocessableEntity)
	})

	t.Run("Create new game with failing insert in store", func(t *testing.T) {
		c, _, _ := testUtils.NewPostJSONRequest(mockGameReqs[0])
		gs.ReturnError = errStore

		err := h.CreateGame(c)
		he := err.(*echo.HTTPError)

		testUtils.AssertHTTPStatus(t, he.Code, http.StatusInternalServerError)
		testUtils.AssertEqual(t, he.Message, gs.ReturnError.Error())
	})

}

func TestGetGames(t *testing.T) {
	setup()

	var activeGames []model.Game
	for _, g := range testUtils.DB1.Games {
		if g.Status == store.GameStatusActive {
			activeGames = append(activeGames, g)
		}
	}

	uid := testUtils.DB1.Users[0].UserID

	t.Run("Get games successfully", func(t *testing.T) {
		c, _, res := testUtils.NewGetRequest()

		c.Set(user.UserIDCtxKey, uid)

		h.GetGames(c)
		testUtils.AssertHTTPStatus(t, res.Code, http.StatusOK)

		var gotGames []model.Game
		json.NewDecoder(res.Body).Decode(&gotGames)

		testUtils.AssertEqual(t, gotGames, testUtils.DB1.Games)
	})

	t.Run("Get active games successfully", func(t *testing.T) {
		c, _, res := testUtils.NewGetRequest()

		c.Set(user.UserIDCtxKey, user.UserIDLimitedAccess)

		h.GetGames(c)
		testUtils.AssertHTTPStatus(t, res.Code, http.StatusOK)

		var gotGames []model.Game
		json.NewDecoder(res.Body).Decode(&gotGames)

		testUtils.AssertEqual(t, gotGames, activeGames)
	})

	t.Run("Get games successfully", func(t *testing.T) {
		c, _, res := testUtils.NewGetRequest()

		c.Set(user.UserIDCtxKey, uid)

		h.GetGames(c)
		testUtils.AssertHTTPStatus(t, res.Code, http.StatusOK)

		var gotGames []model.Game
		json.NewDecoder(res.Body).Decode(&gotGames)

		testUtils.AssertEqual(t, gotGames, testUtils.DB1.Games)
	})

	t.Run("Get games with failing select in store", func(t *testing.T) {
		gs.ReturnError = errStore

		c, _, _ := testUtils.NewGetRequest()

		err := h.GetGames(c)
		he := err.(*echo.HTTPError)

		testUtils.AssertHTTPStatus(t, he.Code, http.StatusInternalServerError)
		testUtils.AssertEqual(t, he.Message, gs.ReturnError.Error())
	})
}
