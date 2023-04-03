package logger

import (
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"strings"

	"github.com/labstack/echo/v4"
	gommonLog "github.com/labstack/gommon/log"
	echoLogrusMidd "github.com/neko-neko/echo-logrus/v2"
	"github.com/neko-neko/echo-logrus/v2/log"
	"github.com/sirupsen/logrus"
)

//Logger logger
var Logger *logrus.Logger

//EchoLogger logger
var EchoLogger *log.MyLogger

//EchoMiddleware returns a middleware that logs Echo HTTP requests.
var EchoMiddleware echo.MiddlewareFunc

func init() {
	// Logger
	log.Logger().SetOutput(os.Stdout)
	log.Logger().SetLevel(gommonLog.INFO)
	log.Logger().SetFormatter(&logrus.TextFormatter{
		FullTimestamp:          true,
		DisableSorting:         true,
		QuoteEmptyFields:       true,
		DisableLevelTruncation: true,
	})

	EchoLogger = log.Logger()

	Logger = logrus.New()

	//disable logger on test runs
	if strings.HasSuffix(os.Args[0], ".test") {
		EchoLogger.SetOutput(ioutil.Discard)
		Logger.SetOutput(ioutil.Discard)
	}

	EchoMiddleware = echoLogrusMidd.Logger()

	Logger.SetLevel(logrus.InfoLevel)
	Logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:          true,
		DisableSorting:         true,
		QuoteEmptyFields:       true,
		DisableLevelTruncation: true,
	})
}

//EchoHTTPError logging error and creates a new echo HTTPError instance
func EchoHTTPError(httpStatusCode int, message ...interface{}) *echo.HTTPError {

	// Print name of caller and line called
	nm, ln := runFuncName()
	tempLog := Logger.WithFields(logrus.Fields{"caller": nm, "line": ln})

	if httpStatusCode >= http.StatusInternalServerError {
		//In case of Internal error override response error message
		tempLog.Error(message)
		return echo.NewHTTPError(httpStatusCode, "internal server error")
	} else {
		tempLog.Warning(message)
		return echo.NewHTTPError(httpStatusCode, message)
	}
}

func runFuncName() (name string, line int) {
	pc := make([]uintptr, 1)
	runtime.Callers(3, pc)
	f := runtime.FuncForPC(pc[0])
	name = f.Name()
	_, line = f.FileLine(pc[0])

	return
}
