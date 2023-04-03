package user

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"gitlab.com/khakibee/khakibee/api/logger"
	"gitlab.com/khakibee/khakibee/api/model"
	"gitlab.com/khakibee/khakibee/api/store"
)

type UserHandler struct {
	store     store.UserStore
	jwtSecret string
}

type AuthUserJWTClaims struct {
	*jwt.StandardClaims
	UserID int `json:"uid"`
}

func NewUserHandler(store store.UserStore, jwtSecret string) (userHandler *UserHandler) {
	userHandler = &UserHandler{store, jwtSecret}
	return
}

func (u UserHandler) Register(api *echo.Group) {
	api.POST("/auth", u.AuthUser)
}

func (u UserHandler) AuthUser(c echo.Context) error {
	var us model.User
	req := new(authUserReq)

	if err := req.bind(c, &us); err != nil {
		return logger.EchoHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	uid, err := u.store.SelectUserByCred(c.Request().Context(), us.Usrnme, us.Psswrd)
	if err != nil {
		return logger.EchoHTTPError(http.StatusInternalServerError, err.Error())
	}
	if uid == nil {
		return logger.EchoHTTPError(http.StatusUnauthorized, "wrong credentials")
	}

	token, err := generateAuthToken(*uid, u.jwtSecret)
	if err != nil {
		return logger.EchoHTTPError(http.StatusInternalServerError, err.Error())
	}
	resp := model.Auth{Token: token}

	return c.JSON(http.StatusOK, resp)
}
