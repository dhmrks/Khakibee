package store_test

import (
	"context"
	"database/sql"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/lib/pq"
	"gitlab.com/khakibee/khakibee/api/model"
	"gitlab.com/khakibee/khakibee/api/store"
	"gitlab.com/khakibee/khakibee/api/testUtils"
)

var selectBkngCols = []string{"bkng_id", "dte", "gme_id", "status", "nme", "mob_num", "email_addr"}
var selectBkngsCols = []string{"bkng_id", "dte", "nme"}

func TestCancelBooking(t *testing.T) {
	db, mock := testUtils.NewMockDB(t)
	s := store.NewPSQLBookingStore(db)
	query := regexp.QuoteMeta(`update bkng set status = 'canceled', canceled_at = current_timestamp	where gme_id=$1 and bkng_id = $2 and status='booked';`)
	mb := testUtils.DB1.Bookings

	t.Run("cancel booking successfully", func(t *testing.T) {
		mock.ExpectExec(query).WithArgs(mb[0].GameID, mb[0].BkngID).
			WillReturnResult(sqlmock.NewResult(0, 1))

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		ok, err := s.CancelBooking(ctx, mb[0].GameID, mb[0].BkngID)

		testUtils.AssertError(t, err, false)
		testUtils.AssertEqual(t, ok, true)
	})

	t.Run("cancel booking not found", func(t *testing.T) {
		mock.ExpectExec(query).WithArgs(mb[0].GameID, mb[0].BkngID).
			WillReturnResult(sqlmock.NewResult(0, 0))

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		ok, err := s.CancelBooking(ctx, mb[0].GameID, mb[0].BkngID)

		testUtils.AssertError(t, err, false)
		testUtils.AssertEqual(t, ok, false)
	})

}

func TestUpdateBooking(t *testing.T) {
	db, mock := testUtils.NewMockDB(t)
	s := store.NewPSQLBookingStore(db)
	query := regexp.QuoteMeta(`update bkng set dte = $3, status = $4, nme = $5, mob_num = $6, email_addr = $7, notes = $8, lst_edted_by = $9 where gme_id = $1 and bkng_id = $2;`)
	mb := testUtils.DB1.Bookings
	uid := testUtils.DB1.Users[0].UserID

	t.Run("update booking successfully", func(t *testing.T) {
		mock.ExpectExec(query).WithArgs(mb[0].BkngID, mb[0].GameID, mb[0].Dte, mb[0].Status, mb[0].Name, mb[0].MobNum, mb[0].Email, mb[0].Notes, uid).
			WillReturnResult(sqlmock.NewResult(0, 1))

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		notFound, dateConflict, err := s.UpdateBooking(ctx, mb[0], uid)

		testUtils.AssertError(t, err, false)
		testUtils.AssertEqual(t, notFound, false)
		testUtils.AssertEqual(t, dateConflict, false)
	})

	t.Run("update booking with date conflict", func(t *testing.T) {
		mock.ExpectExec(query).WithArgs(mb[0].BkngID, mb[0].GameID, mb[0].Dte, mb[0].Status, mb[0].Name, mb[0].MobNum, mb[0].Email, mb[0].Notes, uid).
			WillReturnError(&pq.Error{Code: pq.ErrorCode(store.PQ_UNIQUE_VIOLATION)})

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		notFound, dateConflict, err := s.UpdateBooking(ctx, mb[0], uid)

		testUtils.AssertError(t, err, false)
		testUtils.AssertEqual(t, notFound, false)
		testUtils.AssertEqual(t, dateConflict, true)
	})

	t.Run("update booking with date conflict", func(t *testing.T) {
		mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(0, 0))

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		notFound, dateConflict, err := s.UpdateBooking(ctx, mb[0], uid)

		testUtils.AssertError(t, err, false)
		testUtils.AssertEqual(t, notFound, true)
		testUtils.AssertEqual(t, dateConflict, false)
	})

}

func TestSelectBookingByID(t *testing.T) {
	db, mock := testUtils.NewMockDB(t)
	store := store.NewPSQLBookingStore(db)
	query := regexp.QuoteMeta(`select bkng_id, TO_CHAR(dte, 'YYYY-MM-DD HH24:MI'), gme_id, status, nme, mob_num, email_addr from bkng where gme_id=$1 and bkng_id = $2;`)

	t.Run("select bookings succesfully", func(t *testing.T) {
		mb := testUtils.DB1.Bookings
		rows := sqlmock.NewRows(selectBkngCols).
			AddRow(mb[0].BkngID, mb[0].Dte, mb[0].GameID, mb[0].Status, mb[0].Name, mb[0].MobNum, mb[0].Email)

		mock.ExpectQuery(query).WillReturnRows(rows)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		bkngs, err := store.SelectBookingByID(ctx, mb[0].GameID, mb[0].BkngID)

		testUtils.AssertError(t, err, false)
		testUtils.AssertEqual(t, *bkngs, mb[0])
	})

	t.Run("select booking not found", func(t *testing.T) {

		mock.ExpectQuery(query).WillReturnError(sql.ErrNoRows)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		bkngs, err := store.SelectBookingByID(ctx, 1, 3)

		testUtils.AssertError(t, err, false)
		testUtils.AssertNil(t, bkngs)
	})
}

func TestSelectBookings(t *testing.T) {
	db, mock := testUtils.NewMockDB(t)
	store := store.NewPSQLBookingStore(db)
	query := regexp.QuoteMeta(`select bkng_id, TO_CHAR(dte at time zone $2, 'YYYY-MM-DD HH24:MI'), nme
						from bkng
						where extract(MONTH FROM (CURRENT_TIMESTAMP AT time zone $2) +  $3::interval ) = extract(MONTH FROM dte)
						and canceled_at IS NULL
						and gme_id = $1`)

	tb := testUtils.DB1.Bookings

	t.Run("select bookings succesfully", func(t *testing.T) {
		mb := []model.Booking{
			{
				BkngID: tb[0].BkngID,
				Dte:    tb[0].Dte,
				Name:   tb[0].Name},
			{
				BkngID: tb[1].BkngID,
				Dte:    tb[1].Dte,
				Name:   tb[1].Name},
		}

		rows := sqlmock.NewRows(selectBkngsCols).
			AddRow(mb[0].BkngID, mb[0].Dte, mb[0].Name).
			AddRow(mb[1].BkngID, mb[1].Dte, mb[1].Name)

		mock.ExpectQuery(query).WillReturnRows(rows)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		bkngs, err := store.SelectBookings(ctx, 1, "athens", 0)

		testUtils.AssertError(t, err, false)
		testUtils.AssertEqual(t, bkngs, mb)
	})

	t.Run("select bookings with nil values", func(t *testing.T) {
		mb := testUtils.DB1.Bookings
		rows := sqlmock.NewRows(selectBkngsCols).
			AddRow(nil, mb[0].Dte, mb[0].Name)

		mock.ExpectQuery(query).WillReturnRows(rows)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		bkngs, err := store.SelectBookings(ctx, 1, "Athens", 0)

		testUtils.AssertError(t, err, true)
		testUtils.AssertNil(t, bkngs)
	})
}

func TestInsertBooking(t *testing.T) {
	db, mock := testUtils.NewMockDB(t)
	store := store.NewPSQLBookingStore(db)
	query := regexp.QuoteMeta(`
		insert into bkng (gme_id, dte, nme, mob_num, email_addr, notes, bked_by)
		values ($1, $2, $3, $4, $5, $6, $7)
		ON CONFLICT DO NOTHING
		RETURNING bkng_id;
	`)

	bkng := testUtils.DB1.Bookings[0]
	uid := testUtils.DB1.Users[0].UserID
	wantBkngID := 1

	t.Run("insert booking succesfully", func(t *testing.T) {

		rows := sqlmock.NewRows([]string{"bkng_id"}).
			AddRow(wantBkngID)

		mock.ExpectQuery(query).WillReturnRows(rows)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		goBkngID, err := store.InsertBooking(ctx, bkng, uid)

		testUtils.AssertError(t, err, false)
		testUtils.AssertEqual(t, *goBkngID, wantBkngID)
	})

	t.Run("insert booking duplicated", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"bkng_id"})

		mock.ExpectQuery(query).WillReturnRows(rows)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		gotBkngID, err := store.InsertBooking(ctx, bkng, uid)

		testUtils.AssertError(t, err, false)
		testUtils.AssertNil(t, gotBkngID)
	})
}
