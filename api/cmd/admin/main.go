package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"

	"gitlab.com/khakibee/khakibee/api/config"
	"gitlab.com/khakibee/khakibee/api/email"
	"gitlab.com/khakibee/khakibee/api/game"
	"gitlab.com/khakibee/khakibee/api/logger"
	"gitlab.com/khakibee/khakibee/api/store"
	"gitlab.com/khakibee/khakibee/api/user"
)

type Validator struct {
	validator *validator.Validate
}

const GracefullyShuttingMsg = "Gracefully shutting down server"

func main() {
	cnfg := config.LoadEnv()

	e := echo.New()

	e.Logger = logger.EchoLogger

	e.Validator = &Validator{validator: validator.New()}
	e.Use(middleware.Recover())

	db := connectDB(cnfg.SQLString)
	defer db.Close()

	api := e.Group("/api")

	e.Use(logger.EchoMiddleware)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: cnfg.AllowOrigins,
	}))

	api.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:         &user.AuthUserJWTClaims{},
		SigningKey:     []byte(cnfg.JWTSecret),
		SuccessHandler: user.ParseJWTFromCTX,
		Skipper:        AuthWhitelistSkipper,
	}))

	e.Pre(middleware.RemoveTrailingSlash())

	gameStore := store.NewPSQLGameStore(db)
	bkngStore := store.NewPSQLBookingStore(db)
	userStore := store.NewPSQLUserStore(db)
	schdlStore := store.NewPSQLScheduleStore(db)

	emailStore := store.NewPSQLEmailStore(db)
	emailHandler, err := email.NewSMTPEmailHandler(emailStore, cnfg.SMTPDtls)
	if err != nil {
		log.Fatalf("error with SMTP configuration: %v\n", err)
	}

	gameHandler := game.NewGameHandler(gameStore, bkngStore, schdlStore, emailHandler)
	gameHandler.Register(api)

	userHandler := user.NewUserHandler(userStore, cnfg.JWTSecret)
	userHandler.Register(api)

	// create new dp listener
	n := store.NewListener(cnfg.SQLString)

	//listen for bookings change on db
	go n.ListenAndNotify(store.BkngChangesPSQLChan, gameHandler.DispatchBookingUpdate)

	// Start server
	go func() {
		if err := e.Start(fmt.Sprintf(":%s", cnfg.Port)); err != nil {
			e.Logger.Fatal(err.Error())
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	osSig := make(chan os.Signal, 1)
	signal.Notify(osSig, os.Interrupt, syscall.SIGTERM)
	<-osSig
	e.Logger.Info(GracefullyShuttingMsg)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

}

func (cv *Validator) Validate(i interface{}) error {

	if err := cv.validator.Struct(i); err != nil {
		return err
	}
	return nil
}

func connectDB(connStr string) *sql.DB {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Check if connection is ok
	err = db.Ping()
	if err != nil {
		log.Fatalf("error pinging database: %v\n", err)
	}

	return db
}

// AuthWhitelistSkipper : Skipping Authentication middleware for selected cases
func AuthWhitelistSkipper(c echo.Context) bool {
	r := c.Request()
	if r.Method == http.MethodOptions {
		return true
	}

	allowlist := []string{"/api/auth"}

	for _, al := range allowlist {
		if r.URL.Path == al {
			return true
		}
	}

	return false
}
