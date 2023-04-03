package store

import (
	"context"
	"database/sql"

	"github.com/neko-neko/echo-logrus/v2/log"
	"github.com/sirupsen/logrus"
	"gitlab.com/khakibee/khakibee/api/model"
)

type EmailStore interface {
	SelectEmailTemplate(ctx context.Context, code string) (emailTmpl model.EmailTmpl, err error)
	SelectEmailDtls(ctx context.Context, bkngID int, timezone string) (emailTmpl model.EmailDtls, err error)
	InsertEmailReport(ctx context.Context, bkngID int, code string, emaiErr error) (err error)
}

type PSQLEmailStore struct {
	db *sql.DB
}

func NewPSQLEmailStore(db *sql.DB) (EmailStore *PSQLEmailStore) {
	EmailStore = &PSQLEmailStore{db}

	return
}

func (e *PSQLEmailStore) InsertEmailReport(ctx context.Context, bkngID int, code string, emaiErr error) (err error) {
	query := `insert into mail_rprt (bkng_id, code, error) values ($1, $2, $3);`
	var emailErrorStr *string
	if emaiErr != nil {
		emailErrorStr = new(string)
		*emailErrorStr = emaiErr.Error()
	}

	args := []interface{}{bkngID, code, emailErrorStr}

	_, err = e.db.ExecContext(ctx, query, args...)
	if err != nil {
		log.Logger().WithFields(logrus.Fields{"query": query, "args": args}).Error(err)
	}

	return
}

func (e *PSQLEmailStore) SelectEmailTemplate(ctx context.Context, code string) (emailTmpl model.EmailTmpl, err error) {
	query := `select subject, body 
	from mail_tmpl et
	where et.code = $1;`

	args := []interface{}{code}

	row := e.db.QueryRowContext(ctx, query, args...)
	err = row.Scan(&emailTmpl.Subject, &emailTmpl.Body)
	if err != nil {
		log.Logger().WithFields(logrus.Fields{"query": query, "args": args}).Error(err)
	}

	return
}

func (e *PSQLEmailStore) SelectEmailDtls(ctx context.Context, bkngID int, timezone string) (emailDtls model.EmailDtls, err error) {

	query := `select g.nme, b.nme, b.email_addr, b.mob_num, TO_CHAR(b.dte at time zone $2, 'YYYY-MM-DD HH24:MI'), b.notes
	from bkng b 
	join gme g USING (gme_id)
	where bkng_id = $1;`

	args := []interface{}{bkngID, timezone}

	row := e.db.QueryRowContext(ctx, query, args...)
	err = row.Scan(
		&emailDtls.GameName,
		&emailDtls.Name,
		&emailDtls.Email,
		&emailDtls.MobNum,
		&emailDtls.Date,
		&emailDtls.Notes,
	)
	if err != nil {
		log.Logger().WithFields(logrus.Fields{"query": query, "args": args}).Error(err)
	}

	return
}
