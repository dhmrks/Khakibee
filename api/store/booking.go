package store

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/lib/pq"
	"github.com/neko-neko/echo-logrus/v2/log"
	"github.com/sirupsen/logrus"
	"gitlab.com/khakibee/khakibee/api/model"
)

const PQ_UNIQUE_VIOLATION = "23505"

type BookingStore interface {
	InsertBooking(ctx context.Context, bkng model.Booking, bkedBy int) (bkngID *int, err error)
	SelectBookings(ctx context.Context, gmeID int, timezone string, offset int) (bkngs []model.Booking, err error)
	SelectBookingByID(ctx context.Context, gmeID int, bkngID int) (bkng *model.Booking, err error)
	UpdateBooking(ctx context.Context, bkng model.Booking, lstEditedBy int) (notFound, dateConflict bool, err error)
	CancelBooking(ctx context.Context, gmeID int, bkngID int) (ok bool, err error)
}

type PSQLBookingStore struct {
	db *sql.DB
}

func NewPSQLBookingStore(db *sql.DB) (bookingStore *PSQLBookingStore) {
	bookingStore = &PSQLBookingStore{db}

	return
}

func (g *PSQLBookingStore) SelectBookings(ctx context.Context, gmeID int, timezone string, offset int) (bkngs []model.Booking, err error) {
	query := `select bkng_id, TO_CHAR(dte at time zone $2, 'YYYY-MM-DD HH24:MI'), nme
						from bkng
						where extract(MONTH FROM (CURRENT_TIMESTAMP AT time zone $2) +  $3::interval ) = extract(MONTH FROM dte)
						and canceled_at IS NULL
						and gme_id = $1;`

	interval := fmt.Sprintf("%d month", offset)
	args := []interface{}{gmeID, timezone, interval}

	rows, err := g.db.QueryContext(ctx, query, args...)
	if err != nil {
		log.Logger().WithFields(logrus.Fields{"query": query}).Error(err)
		return
	}

	for rows.Next() {
		var b model.Booking
		err = rows.Scan(
			&b.BkngID,
			&b.Dte,
			&b.Name,
		)

		if err != nil {
			return nil, err
		}
		bkngs = append(bkngs, b)
	}

	return
}

func (g *PSQLBookingStore) SelectBookingByID(ctx context.Context, gmeID int, bkngID int) (bkng *model.Booking, err error) {
	query := `select bkng_id, TO_CHAR(dte, 'YYYY-MM-DD HH24:MI'), gme_id, status, nme, mob_num, email_addr from bkng where gme_id=$1 and bkng_id = $2;`
	args := []interface{}{gmeID, bkngID}

	row := g.db.QueryRowContext(ctx, query, args...)
	var b model.Booking
	err = row.Scan(
		&b.BkngID,
		&b.Dte,
		&b.GameID,
		&b.Status,
		&b.Name,
		&b.MobNum,
		&b.Email,
	)
	if err == sql.ErrNoRows {
		err = nil
	} else if err != nil {
		log.Logger().WithFields(logrus.Fields{"query": query}).Error(err)
	} else {
		bkng = &b
	}

	return
}

func (g *PSQLBookingStore) UpdateBooking(ctx context.Context, bkng model.Booking, lstEditedBy int) (notFound, dateConflict bool, err error) {
	query := `update bkng set dte = $3, status = $4, nme = $5, mob_num = $6, email_addr = $7, notes = $8, lst_edted_by = $9 where gme_id = $1 and bkng_id = $2;`

	fmt.Println("UpdateBooking: ", query, bkng.Status)
	args := []interface{}{bkng.GameID, bkng.BkngID, bkng.Dte, bkng.Status, bkng.Name, bkng.MobNum, bkng.Email, bkng.Notes, lstEditedBy}

	res, err := g.db.ExecContext(ctx, query, args...)
	if err != nil {

		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == pq.ErrorCode(PQ_UNIQUE_VIOLATION) {
			err = nil
			dateConflict = true
			return
		}

		log.Logger().WithFields(logrus.Fields{"query": query, "args": args}).Error(err)
		return
	}

	ra, err := res.RowsAffected()
	notFound = ra == 0

	return
}

func (g *PSQLBookingStore) CancelBooking(ctx context.Context, gmeID int, bkngID int) (ok bool, err error) {
	query := `
		update bkng set status = 'canceled', canceled_at = current_timestamp
		where gme_id=$1 and bkng_id = $2 and status='booked';
	`

	args := []interface{}{gmeID, bkngID}

	res, err := g.db.ExecContext(ctx, query, args...)
	if err != nil {
		log.Logger().WithFields(logrus.Fields{"query": query, "args": args}).Error(err)
		return
	}

	ra, err := res.RowsAffected()
	ok = ra == 1

	return
}

func (g *PSQLBookingStore) InsertBooking(ctx context.Context, bkng model.Booking, bkedBy int) (bkngID *int, err error) {
	query := `
		insert into bkng (gme_id, dte, nme, mob_num, email_addr, notes, bked_by)
		values ($1, $2, $3, $4, $5, $6, $7)
		ON CONFLICT DO NOTHING
		RETURNING bkng_id;
	`

	args := []interface{}{bkng.GameID, bkng.Dte, bkng.Name, bkng.MobNum, bkng.Email, bkng.Notes, bkedBy}

	row := g.db.QueryRowContext(ctx, query, args...)

	err = row.Scan(&bkngID)
	if err == sql.ErrNoRows {
		err = nil
	} else if err != nil {
		log.Logger().WithFields(logrus.Fields{"query": query, "args": args}).Error(err)
		return
	}

	return
}
