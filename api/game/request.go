package game

import (
	"time"

	"github.com/labstack/echo/v4"
	"gitlab.com/khakibee/khakibee/api/cmd/admin/common"
	"gitlab.com/khakibee/khakibee/api/model"
)

type putNewScheduleReq struct {
	GameID int `param:"gid"`
	model.WeekSchedule
	ActiveBy string `json:"active_by" validate:"required,datetime=2006-01-02"`
}

func (r *putNewScheduleReq) bind(c echo.Context, gmeID *int, s *model.WeekSchedule, activeBy *string) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}

	*gmeID = r.GameID
	*s = r.WeekSchedule
	*activeBy = r.ActiveBy

	return nil
}

type putScheduleReq struct {
	GameID int `param:"gid"`
	model.WeekSchedule
}

func (r *putScheduleReq) bind(c echo.Context, gmeID *int, s *model.WeekSchedule) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}

	*gmeID = r.GameID
	*s = r.WeekSchedule

	return nil
}

type getScheduleReq struct {
	GameID int `param:"gid"`
}

func (r *getScheduleReq) bind(c echo.Context, gmeID *int) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}

	*gmeID = r.GameID

	return nil
}

type removeBookingReq struct {
	GameID int `param:"gid"`
	BkngID int `param:"bid"`
}

func (r *removeBookingReq) bind(c echo.Context, b *model.Booking, location *time.Location) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}

	b.GameID = r.GameID
	b.BkngID = r.BkngID
	*location = common.GetHeaderLocation(c)

	return nil
}

type getBookingReq struct {
	GameID int `param:"gid"`
	BkngID int `param:"bid"`
}

func (r *getBookingReq) bind(c echo.Context, b *model.Booking) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}

	b.GameID = r.GameID
	b.BkngID = r.BkngID

	return nil
}

type getBookingsReq struct {
	GameID int `param:"gid"`
}

func (r *getBookingsReq) bind(c echo.Context, b *model.Booking, location *time.Location) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}

	b.GameID = r.GameID
	*location = common.GetHeaderLocation(c)

	return nil
}

type getCalendarReq struct {
	GameID int `param:"gid"`

	Offset int `query:"offset"`
}

func (r *getCalendarReq) bind(c echo.Context, g *model.Game, offset *int, location *time.Location) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}

	g.GameID = r.GameID
	*offset = r.Offset
	*location = common.GetHeaderLocation(c)

	return nil
}

type getGameReq struct {
	GameID int `param:"gid"`
}

func (r *getGameReq) bind(c echo.Context, b *model.Game) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}

	b.GameID = r.GameID

	return nil
}

type createGameReq struct {
	Name  string `json:"name"  validate:"required"`
	Descr string `json:"descr"  validate:"required"`
	Addr  string `json:"addr" validate:"required"`

	Status string `json:"status" validate:"required"`

	ImgURL   string `json:"img_url" validate:"required,uri"`
	MapURL   string `json:"map_url" validate:"required,uri"`
	Plrs     string `json:"players" validate:"required"`
	Dur      int    `json:"duration" validate:"required"`
	AgeRange string `json:"age_range" validate:"required"`
}

func (r *createGameReq) bind(c echo.Context, g *model.Game) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}

	g.Name = r.Name
	g.Descr = r.Descr
	g.Addr = r.Addr
	g.Status = r.Status
	g.ImgURL = r.ImgURL
	g.MapURL = r.MapURL
	g.Plrs = r.Plrs
	g.Dur = r.Dur
	g.AgeRange = r.AgeRange

	return nil
}

type editGameReq struct {
	GameID int    `param:"gid"`
	Status string `json:"status" validate:"required"`

	Name  string `json:"name"  validate:"required"`
	Descr string `json:"descr"  validate:"required"`
	Addr  string `json:"addr" validate:"required"`

	ImgURL   string `json:"img_url" validate:"required,uri"`
	MapURL   string `json:"map_url" validate:"required,uri"`
	Plrs     string `json:"players" validate:"required"`
	Dur      int    `json:"duration" validate:"required"`
	AgeRange string `json:"age_range" validate:"required"`
}

func (r *editGameReq) bind(c echo.Context, g *model.Game) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}

	g.GameID = r.GameID
	g.Status = r.Status
	g.Name = r.Name
	g.Descr = r.Descr
	g.Addr = r.Addr
	g.ImgURL = r.ImgURL
	g.MapURL = r.MapURL
	g.Plrs = r.Plrs
	g.Dur = r.Dur
	g.AgeRange = r.AgeRange

	return nil
}

type bookGameReq struct {
	Dte    string `json:"date"  validate:"required,datetime=2006-01-02 15:04"`
	GameID int    `param:"gid"`

	Name   string `json:"name"  validate:"required"`
	MobNum string `json:"mob_num"  validate:"required,e164"`
	Email  string `json:"email" validate:"required,email"`

	Notes *string `json:"notes"`
}

func (r *bookGameReq) bind(c echo.Context, b *model.Booking, location *time.Location) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}

	b.Dte = r.Dte
	b.GameID = r.GameID
	b.Name = r.Name
	b.MobNum = r.MobNum
	b.Email = r.Email
	b.Notes = r.Notes
	*location = common.GetHeaderLocation(c)

	return nil
}

type editBookingReq struct {
	BkngID int    `param:"bid"`
	Status string `json:"status"  validate:"required"`

	Dte    string `json:"date"  validate:"required,datetime=2006-01-02 15:04"`
	GameID int    `param:"gid"`

	Name   string `json:"name"  validate:"required"`
	MobNum string `json:"mob_num"  validate:"required,e164"`
	Email  string `json:"email" validate:"required,email"`

	Notes *string `json:"notes"`
}

func (r *editBookingReq) bind(c echo.Context, b *model.Booking, location *time.Location) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}

	b.BkngID = r.BkngID
	b.Status = r.Status
	b.Dte = r.Dte
	b.GameID = r.GameID
	b.Name = r.Name
	b.MobNum = r.MobNum
	b.Email = r.Email
	b.Notes = r.Notes
	*location = common.GetHeaderLocation(c)

	return nil
}
