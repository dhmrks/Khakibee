package common

import (
	"time"

	"github.com/labstack/echo/v4"
)

func GetHeaderLocation(c echo.Context) (location time.Location) {

	var (
		timezone string
		loc      *time.Location
		err      error
	)
	if t := c.Request().Header["Timezone"]; len(t) == 1 {
		timezone = t[0]
	}

	if loc, err = time.LoadLocation(timezone); err != nil {
		loc = time.Now().Location()
	}

	return *loc
}
