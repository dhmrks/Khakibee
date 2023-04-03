package user_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
	"gitlab.com/khakibee/khakibee/api/model"
	"gitlab.com/khakibee/khakibee/api/store"
	"gitlab.com/khakibee/khakibee/api/testUtils"
	"gitlab.com/khakibee/khakibee/api/user"
)

var mockUsers = []model.User{
	{
		UserID: 1,
		Usrnme: "ioulios",
		Psswrd: "pass",
	},
}

var (
	s        *store.MockUserStore
	h        *user.UserHandler
	errStore = errors.New("internal server error")
)

func setup() {
	s = &store.MockUserStore{Users: testUtils.DB1.Users}
	h = user.NewUserHandler(s, jwtSecret)
}

const jwtSecret = "secret"

var credentials = []string{
	`{"username":"ioulios","password":"pass"}`,
	`{"username":"wrongusername","password":"pass"}`,
}

func TestAuthUser(t *testing.T) {
	setup()

	t.Run("authenticate user successfully", func(t *testing.T) {
		c, _, res := testUtils.NewPostJSONRequest(credentials[0])

		err := h.AuthUser(c)

		testUtils.AssertError(t, err, false)
		testUtils.AssertHTTPStatus(t, res.Code, http.StatusOK)

		ar := &model.Auth{}
		json.NewDecoder(res.Body).Decode(&ar)

		testUtils.AssertAuthJWT(t, ar.Token, jwtSecret, user.AuthUserJWTClaims{UserID: 1})
	})

	t.Run("authenticate user with wrong credentials", func(t *testing.T) {
		c, _, _ := testUtils.NewPostJSONRequest(credentials[1])

		err := h.AuthUser(c).(*echo.HTTPError)

		testUtils.AssertError(t, err, true)
		testUtils.AssertHTTPStatus(t, err.Code, http.StatusUnauthorized)
	})

	t.Run("authenticate user with invalid json data", func(t *testing.T) {
		c, _, _ := testUtils.NewPostJSONRequest(`{}`)

		err := h.AuthUser(c)
		he := err.(*echo.HTTPError)

		testUtils.AssertHTTPStatus(t, he.Code, http.StatusUnprocessableEntity)
		testUtils.AssertError(t, he, true)
	})

	t.Run("authenticate user with store error", func(t *testing.T) {
		s.ReturnError = errStore
		c, _, _ := testUtils.NewPostJSONRequest(credentials[0])

		err := h.AuthUser(c).(*echo.HTTPError)

		testUtils.AssertHTTPStatus(t, err.Code, http.StatusInternalServerError)
		testUtils.AssertEqual(t, err.Message, s.ReturnError.Error())
	})
}
