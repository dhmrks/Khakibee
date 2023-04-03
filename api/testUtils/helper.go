package testUtils

import (
	"database/sql"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

const (
	HeaderTimezone = "Timezone"
	Timezone       = "Europe/Athens"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return err
	}
	return nil
}

func NewDeleteRequest() (c echo.Context, request *http.Request, response *httptest.ResponseRecorder) {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	request, _ = http.NewRequest(http.MethodDelete, "/", nil)
	request.Header.Set(HeaderTimezone, Timezone)

	response = httptest.NewRecorder()
	c = e.NewContext(request, response)

	return
}

func NewGetRequest() (c echo.Context, request *http.Request, response *httptest.ResponseRecorder) {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	request, _ = http.NewRequest(http.MethodGet, "/", nil)
	request.Header.Set(HeaderTimezone, Timezone)

	response = httptest.NewRecorder()
	c = e.NewContext(request, response)

	return
}

func NewPostJSONRequest(body string) (c echo.Context, request *http.Request, response *httptest.ResponseRecorder) {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	request, _ = http.NewRequest(http.MethodPost, "/", strings.NewReader(body))
	request.Header.Set(HeaderTimezone, Timezone)
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	response = httptest.NewRecorder()
	c = e.NewContext(request, response)

	return
}

func NewPutJSONRequest(body string) (c echo.Context, request *http.Request, response *httptest.ResponseRecorder) {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	request, _ = http.NewRequest(http.MethodPut, "/", strings.NewReader(body))
	request.Header.Set(HeaderTimezone, Timezone)
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	response = httptest.NewRecorder()
	c = e.NewContext(request, response)

	return
}

func NewMockDB(t testing.TB) (*sql.DB, sqlmock.Sqlmock) {
	t.Helper()
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}
