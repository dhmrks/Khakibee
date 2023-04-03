package store_test

import (
	"context"
	"database/sql"
	"encoding/json"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"gitlab.com/khakibee/khakibee/api/store"
	"gitlab.com/khakibee/khakibee/api/testUtils"
)

var selectScheduleCols = []string{"schdl_id", "schedule", "dur", "active_by"}

func TestUpsertNewSchedule(t *testing.T) {
	db, mock := testUtils.NewMockDB(t)
	s := store.NewPSQLScheduleStore(db)

	queryIns := `insert into schdl (gme_id, schedule,active_by) values ($1,$2,$3)`
	queryUPd := `update schdl set schedule = $1, active_by=$2 where schdl_id=$3`

	mb := testUtils.DB1.Bookings[0]
	ms := *testUtils.DB1.Schedule[mb.GameID][0]
	msByte, _ := json.Marshal(ms.WkSchdl)

	t.Run("insert upcoming schedule succesfully", func(t *testing.T) {
		mock.ExpectExec(queryIns).WithArgs(mb.GameID, msByte, ms.ActiveBy).
			WillReturnResult(sqlmock.NewResult(0, 1))

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err := s.UpsertNewSchedule(ctx, mb.GameID, ms.WkSchdl, ms.ActiveBy, nil)
		testUtils.AssertError(t, err, false)
	})

	t.Run("update upcoming schedule succesfully", func(t *testing.T) {
		mock.ExpectExec(queryUPd).WithArgs(msByte, ms.ActiveBy, ms.SchdlID).
			WillReturnResult(sqlmock.NewResult(0, 1))

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err := s.UpsertNewSchedule(ctx, mb.GameID, ms.WkSchdl, ms.ActiveBy, &ms.SchdlID)
		testUtils.AssertError(t, err, false)
	})
}

func TestUpsertSchedule(t *testing.T) {
	db, mock := testUtils.NewMockDB(t)
	s := store.NewPSQLScheduleStore(db)

	queryIns := `insert into schdl (gme_id, schedule, active_by) values ($1,$2,current_date)`
	queryUPd := `update schdl set schedule = $1 where schdl_id=$2`

	mb := testUtils.DB1.Bookings[0]
	ms := *testUtils.DB1.Schedule[mb.GameID][0]
	msByte, _ := json.Marshal(ms.WkSchdl)

	t.Run("insert schedule succesfully", func(t *testing.T) {
		mock.ExpectExec(queryIns).WithArgs(mb.GameID, msByte).
			WillReturnResult(sqlmock.NewResult(0, 1))

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err := s.UpsertSchedule(ctx, mb.GameID, ms.WkSchdl, nil)
		testUtils.AssertError(t, err, false)
	})

	t.Run("update schedule succesfully", func(t *testing.T) {
		mock.ExpectExec(queryUPd).WithArgs(msByte, ms.SchdlID).
			WillReturnResult(sqlmock.NewResult(0, 1))

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err := s.UpsertSchedule(ctx, mb.GameID, ms.WkSchdl, &ms.SchdlID)
		testUtils.AssertError(t, err, false)
	})
}

func TestSelectScheduleByGameID(t *testing.T) {
	db, mock := testUtils.NewMockDB(t)
	s := store.NewPSQLScheduleStore(db)
	query := regexp.QuoteMeta(`Select schdl_id, schedule, dur, TO_CHAR(active_by, 'YYYY-MM-DD')
	from schdl s 
		JOIN gme g ON g.gme_id = s.gme_id 
	where s.gme_id = $1 
	and active_by <= CURRENT_DATE order by active_by DESC limit	1;`)

	queryUpcoming := regexp.QuoteMeta(`Select schdl_id, schedule, dur, TO_CHAR(active_by, 'YYYY-MM-DD')
	from schdl s 
		JOIN gme g ON g.gme_id = s.gme_id 
	where s.gme_id = $1 
	and active_by > CURRENT_DATE;`)

	mb := testUtils.DB1.Bookings[0]
	ms := *testUtils.DB1.Schedule[mb.GameID][0]
	msByte, _ := json.Marshal(ms.WkSchdl)

	t.Run("select schedule succesfully", func(t *testing.T) {
		rows := sqlmock.NewRows(selectScheduleCols).
			AddRow(ms.SchdlID, msByte, ms.Dur, ms.ActiveBy)

		mock.ExpectQuery(query).WillReturnRows(rows)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		schdl, err := s.SelectScheduleByGameID(ctx, mb.GameID, false)

		testUtils.AssertError(t, err, false)
		testUtils.AssertEqual(t, schdl, &ms)
	})

	t.Run("select upcoming succesfully", func(t *testing.T) {
		rows := sqlmock.NewRows(selectScheduleCols).
			AddRow(ms.SchdlID, msByte, ms.Dur, ms.ActiveBy)

		mock.ExpectQuery(queryUpcoming).WillReturnRows(rows)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		schdl, err := s.SelectScheduleByGameID(ctx, mb.GameID, true)

		testUtils.AssertError(t, err, false)
		testUtils.AssertEqual(t, schdl, &ms)
	})

	t.Run("select schedule empty", func(t *testing.T) {
		mock.ExpectQuery(query).WillReturnError(sql.ErrNoRows)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		schdl, err := s.SelectScheduleByGameID(ctx, mb.GameID, false)

		testUtils.AssertError(t, err, false)
		testUtils.AssertNil(t, schdl)
	})

}

func TestCheckSchedule(t *testing.T) {
	db, mock := testUtils.NewMockDB(t)
	s := store.NewPSQLScheduleStore(db)
	query := regexp.QuoteMeta(`select schedule -> extract(dow from $2::date)::TEXT ? $3 as exists 
	from schdl 
	where gme_id = $1 and active_by <= $2::date order by active_by DESC limit 1;`)

	t.Run("check schedule for existing slot successfully", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"exists"}).AddRow(true)
		mock.ExpectQuery(query).WillReturnRows(rows)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		exists, err := s.CheckSchedule(ctx, "2022-04-03 14:00", 1)

		testUtils.AssertError(t, err, false)
		testUtils.AssertEqual(t, exists, true)
	})

	t.Run("check schedule for not existing slot successfully", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"exists"}).AddRow(false)
		mock.ExpectQuery(query).WillReturnRows(rows)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		exists, err := s.CheckSchedule(ctx, "2022-04-03 14:00", 1)

		testUtils.AssertError(t, err, false)
		testUtils.AssertEqual(t, exists, false)
	})
}
