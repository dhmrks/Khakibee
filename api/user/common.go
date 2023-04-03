package user

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	UserIDCtxKey        = "userIDkey"
	UserIDLimitedAccess = -1
)

// generateAuthToken return authentication JWT
func generateAuthToken(userID int, secret string) (signedToken string, err error) {

	// Create the Claims
	claims := &AuthUserJWTClaims{
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 365).Unix(),
		},
		userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err = token.SignedString([]byte(secret))

	return
}

// ParseJWTFromCTX parses jwt claims from context and set values to new context keys
func ParseJWTFromCTX(c echo.Context) {
	userJWT := c.Get(middleware.DefaultJWTConfig.ContextKey)
	claims := userJWT.(*jwt.Token).Claims.(*AuthUserJWTClaims)

	userID := claims.UserID

	c.Set(UserIDCtxKey, userID)
}

func SetLimitedAccessUserToCTX(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set(UserIDCtxKey, UserIDLimitedAccess)
		return next(c)
	}
}
