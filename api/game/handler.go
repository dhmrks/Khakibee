package game

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/khakibee/khakibee/api/email"
	"gitlab.com/khakibee/khakibee/api/store"
)

type GameHandler struct {
	store          store.GameStore
	bkngStore      store.BookingStore
	schdlStore     store.ScheduleStore
	emailHandler   email.EmailHandler
	bookingChanged chan struct{}
}

func NewGameHandler(store store.GameStore, bkngStore store.BookingStore, schdlStore store.ScheduleStore, emailHandler email.EmailHandler) (gameHandler *GameHandler) {
	gameHandler = &GameHandler{
		store,
		bkngStore,
		schdlStore,
		emailHandler,
		make(chan struct{}),
	}

	return
}

func (g *GameHandler) Register(api *echo.Group) {
	gg := api.Group("/games")
	gg.GET("", g.GetGames)
	gg.GET("/:gid", g.GetGame)
	gg.POST("", g.CreateGame)
	gg.PUT("/:gid", g.EditGame)

	gg.GET("/:gid/calendar", g.GetCalendar)
	gg.GET("/:gid/calendar/sse", g.GetCalendarSSE)

	bg := gg.Group("/:gid/bookings")
	bg.POST("", g.BookGame)
	bg.GET("", g.GetBookings)
	bg.GET("/:bid", g.GetBooking)
	bg.PUT("/:bid", g.EditBooking)
	bg.DELETE("/:bid", g.RemoveBooking)

	sg := gg.Group("/:gid/schedule")
	sg.GET("", g.GetSchedule)
	sg.PUT("", g.PutSchedule)
	sg.PUT("/new", g.PutNewSchedule)
}

func (g *GameHandler) DispatchBookingUpdate() {
	select {
	case g.bookingChanged <- struct{}{}:
	default:
	}
}
