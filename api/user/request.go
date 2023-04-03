package user

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/khakibee/khakibee/api/model"
)

type authUserReq struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (r *authUserReq) bind(c echo.Context, u *model.User) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}

	u.Usrnme = r.Username
	u.Psswrd = r.Password

	return nil
}
