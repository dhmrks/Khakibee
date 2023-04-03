package store

import (
	"context"
	"database/sql"

	"gitlab.com/khakibee/khakibee/api/model"
)

type MockEmailStore struct {
	SelectEmailDtlsCalled   int
	SelectEmailTmplCalled   int
	InsertEmailReportCalled int
	EmailDtls               model.EmailDtls
	EmailTmpl               map[string]model.EmailTmpl
	ReturnError             error
}

func NewMockEmailStore(db *sql.DB) (EmailStore *MockEmailStore) {
	EmailStore = &MockEmailStore{}

	return
}

func (e *MockEmailStore) InsertEmailReport(ctx context.Context, bkngID int, code string, emaiErr error) (err error) {
	e.InsertEmailReportCalled += 1
	err = e.ReturnError
	return err
}

func (e *MockEmailStore) SelectEmailTemplate(ctx context.Context, code string) (emailTmpl model.EmailTmpl, err error) {
	e.SelectEmailTmplCalled += 1
	err = e.ReturnError
	return e.EmailTmpl[code], err
}

func (e *MockEmailStore) SelectEmailDtls(ctx context.Context, bkngID int, timezone string) (emailDtls model.EmailDtls, err error) {
	e.SelectEmailDtlsCalled += 1
	err = e.ReturnError
	return e.EmailDtls, err
}
