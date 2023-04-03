package store_test

import (
	"context"
	"database/sql"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"gitlab.com/khakibee/khakibee/api/store"
	"gitlab.com/khakibee/khakibee/api/testUtils"
)

var selectEmailTmpltCols = []string{"subject", "body"}
var selectEmailDtlstCols = []string{"name", "name", "email_addr", "mob_num", "dte", "notes"}

func TestInsertEmailReport(t *testing.T) {
	db, mock := testUtils.NewMockDB(t)
	s := store.NewPSQLEmailStore(db)
	query := regexp.QuoteMeta(`insert into mail_rprt (bkng_id, code, error) values ($1, $2, $3);`)

	t.Run("insert email report succesfully", func(t *testing.T) {
		mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err := s.InsertEmailReport(ctx, 1, "code", nil)

		testUtils.AssertError(t, err, false)

	})
}

func TestSelectEmailTemplate(t *testing.T) {
	db, mock := testUtils.NewMockDB(t)
	s := store.NewPSQLEmailStore(db)
	query := regexp.QuoteMeta(`select subject, body 
	from mail_tmpl et
	where et.code = $1;`)

	me := testUtils.DB1.EmailTmpl["correct"]

	t.Run("select email template succesfully", func(t *testing.T) {
		rows := sqlmock.NewRows(selectEmailTmpltCols).
			AddRow(me.Subject, me.Body)

		mock.ExpectQuery(query).WillReturnRows(rows)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		tmpl, err := s.SelectEmailTemplate(ctx, "code")

		testUtils.AssertError(t, err, false)
		testUtils.AssertEqual(t, tmpl, me)

	})

	t.Run("select email template with no rows", func(t *testing.T) {
		rows := sqlmock.NewRows(selectEmailTmpltCols)

		mock.ExpectQuery(query).WillReturnRows(rows)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_, err := s.SelectEmailTemplate(ctx, "code")

		testUtils.AssertError(t, err, true)
		testUtils.AssertEqual(t, err, sql.ErrNoRows)

	})

}

func TestSelectEmailDtls(t *testing.T) {
	db, mock := testUtils.NewMockDB(t)
	s := store.NewPSQLEmailStore(db)
	query := regexp.QuoteMeta(`select g.nme, b.nme, b.email_addr, b.mob_num, TO_CHAR(b.dte at time zone $2, 'YYYY-MM-DD HH24:MI'), b.notes
	from bkng b 
	join gme g USING (gme_id)
	where bkng_id = $1;`)

	me := testUtils.DB1.EmailDtls

	t.Run("select email details succesfully", func(t *testing.T) {
		rows := sqlmock.NewRows(selectEmailDtlstCols).
			AddRow(me.GameName, me.Name, me.Email, me.MobNum, me.Date, me.Notes)

		mock.ExpectQuery(query).WillReturnRows(rows)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		dtls, err := s.SelectEmailDtls(ctx, 1, "athens")

		testUtils.AssertError(t, err, false)
		testUtils.AssertEqual(t, dtls, me)

	})

	t.Run("select email details with no rows", func(t *testing.T) {
		rows := sqlmock.NewRows(selectEmailDtlstCols)

		mock.ExpectQuery(query).WillReturnRows(rows)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_, err := s.SelectEmailDtls(ctx, 1, "athens")

		testUtils.AssertError(t, err, true)
		testUtils.AssertEqual(t, err, sql.ErrNoRows)
	})

}
