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

var selectGameRows = []string{"gme_id", "status", "img_url", "map_url", "plrs", "dur", "age_rng", "nme", "descr", "addr"}

func TestSelectGames(t *testing.T) {

	db, mock := testUtils.NewMockDB(t)
	s := store.NewPSQLGameStore(db)
	query := regexp.QuoteMeta("select g.gme_id, g.status, g.img_url, g.map_url, g.plrs, g.dur, g.age_rng, g.nme, g.descr, g.addr from gme g")

	t.Run("select games with status successfully", func(t *testing.T) {
		query := regexp.QuoteMeta("select g.gme_id, g.status, g.img_url, g.map_url, g.plrs, g.dur, g.age_rng, g.nme, g.descr, g.addr from gme g where status = $1")

		mg := testUtils.DB1.Games
		rows := sqlmock.NewRows(selectGameRows).
			AddRow(mg[0].GameID, mg[0].Status, mg[0].ImgURL, mg[0].MapURL, mg[0].Plrs, mg[0].Dur, mg[0].AgeRange, mg[0].Name, mg[0].Descr, mg[0].Addr).
			AddRow(mg[1].GameID, mg[1].Status, mg[1].ImgURL, mg[1].MapURL, mg[1].Plrs, mg[1].Dur, mg[1].AgeRange, mg[1].Name, mg[1].Descr, mg[1].Addr).
			AddRow(mg[2].GameID, mg[2].Status, mg[2].ImgURL, mg[2].MapURL, mg[2].Plrs, mg[2].Dur, mg[2].AgeRange, mg[2].Name, mg[2].Descr, mg[2].Addr)

		mock.ExpectQuery(query).WillReturnRows(rows)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		games, err := s.SelectGames(ctx, store.GameStatusActive)

		testUtils.AssertError(t, err, false)
		testUtils.AssertEqual(t, games, mg)
	})

	t.Run("select games successfully", func(t *testing.T) {
		mg := testUtils.DB1.Games
		rows := sqlmock.NewRows(selectGameRows).
			AddRow(mg[0].GameID, mg[0].Status, mg[0].ImgURL, mg[0].MapURL, mg[0].Plrs, mg[0].Dur, mg[0].AgeRange, mg[0].Name, mg[0].Descr, mg[0].Addr).
			AddRow(mg[1].GameID, mg[1].Status, mg[1].ImgURL, mg[1].MapURL, mg[1].Plrs, mg[1].Dur, mg[1].AgeRange, mg[1].Name, mg[1].Descr, mg[1].Addr).
			AddRow(mg[2].GameID, mg[2].Status, mg[2].ImgURL, mg[2].MapURL, mg[2].Plrs, mg[2].Dur, mg[2].AgeRange, mg[2].Name, mg[2].Descr, mg[2].Addr)

		mock.ExpectQuery(query).WillReturnRows(rows)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		games, err := s.SelectGames(ctx, store.GameStatusAll)

		testUtils.AssertError(t, err, false)
		testUtils.AssertEqual(t, games, mg)
	})

	t.Run("select games with no rows", func(t *testing.T) {
		rows := sqlmock.NewRows(selectGameRows)

		mock.ExpectQuery(query).WillReturnRows(rows)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		games, err := s.SelectGames(ctx, store.GameStatusAll)

		testUtils.AssertError(t, err, false)
		testUtils.AssertEqual(t, len(games), 0)
	})

	t.Run("select games with nil values", func(t *testing.T) {
		mg := testUtils.DB1.Games
		rows := sqlmock.NewRows(selectGameRows).
			AddRow(nil, mg[0].Status, mg[0].ImgURL, mg[0].MapURL, mg[0].Plrs, mg[0].Dur, mg[0].AgeRange, mg[0].Name, mg[0].Descr, mg[0].Addr)

		mock.ExpectQuery(query).WillReturnRows(rows)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_, err := s.SelectGames(ctx, store.GameStatusAll)

		testUtils.AssertError(t, err, true)

	})

	t.Run("select games with nil values", func(t *testing.T) {
		rows := sqlmock.NewRows(selectGameRows)

		mock.ExpectQuery(query).WillReturnRows(rows)

		ctx, cancel := context.WithTimeout(context.Background(), 0)
		defer cancel()
		_, err := s.SelectGames(ctx, store.GameStatusAll)

		testUtils.AssertError(t, err, true)

	})

}

func TestSelectGameByID(t *testing.T) {
	db, mock := testUtils.NewMockDB(t)
	s := store.NewPSQLGameStore(db)
	query := regexp.QuoteMeta("select g.gme_id, g.status, g.img_url, g.map_url, g.plrs, g.dur, g.age_rng, g.nme, g.descr, g.addr from gme g where g.gme_id = $1;")
	mg := testUtils.DB1.Games

	t.Run("select game by id successfully", func(t *testing.T) {
		rows := sqlmock.NewRows(selectGameRows).
			AddRow(mg[0].GameID, mg[0].Status, mg[0].ImgURL, mg[0].MapURL, mg[0].Plrs, mg[0].Dur, mg[0].AgeRange, mg[0].Name, mg[0].Descr, mg[0].Addr)

		mock.ExpectQuery(query).WithArgs(mg[0].GameID).WillReturnRows(rows)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		game, err := s.SelectGameByID(ctx, mg[0].GameID)

		testUtils.AssertError(t, err, false)
		testUtils.AssertEqual(t, game, &mg[0])
	})

	t.Run("select game by id with non existent id", func(t *testing.T) {
		mock.ExpectQuery(query).WillReturnError(sql.ErrNoRows)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		game, err := s.SelectGameByID(ctx, 0)

		testUtils.AssertError(t, err, false)
		testUtils.AssertNil(t, game)

	})

	t.Run("select game by id with nil values", func(t *testing.T) {
		rows := sqlmock.NewRows(selectGameRows).
			AddRow(nil, mg[0].Status, mg[0].ImgURL, mg[0].MapURL, mg[0].Plrs, mg[0].Dur, mg[0].AgeRange, mg[0].Name, mg[0].Descr, mg[0].Addr)

		mock.ExpectQuery(query).WillReturnRows(rows)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_, err := s.SelectGameByID(ctx, mg[0].GameID)

		testUtils.AssertError(t, err, true)

	})

	t.Run("select game by id with context deadline exceeded", func(t *testing.T) {
		rows := sqlmock.NewRows(selectGameRows)

		mock.ExpectQuery(query).WillReturnRows(rows)

		ctx, cancel := context.WithTimeout(context.Background(), 0)
		defer cancel()
		_, err := s.SelectGameByID(ctx, 0)

		testUtils.AssertError(t, err, true)

	})

}

func TestInsertGame(t *testing.T) {

	db, mock := testUtils.NewMockDB(t)
	s := store.NewPSQLGameStore(db)
	query := regexp.QuoteMeta("insert into gme (status, img_url, map_url, plrs, dur, age_rng, nme, descr, addr) values ($1, $2, $3, $4, $5, $6, $7, $8, $9);")
	mg := testUtils.DB1.Games

	mock.ExpectExec(query).WithArgs(mg[0].Status, mg[0].ImgURL, mg[0].MapURL, mg[0].Plrs, mg[0].Dur, mg[0].AgeRange, mg[0].Name, mg[0].Descr, mg[0].Addr).
		WillReturnResult(sqlmock.NewResult(0, 1))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := s.InsertGame(ctx, mg[0])

	testUtils.AssertError(t, err, false)

}

func TestUpdateGame(t *testing.T) {

	db, mock := testUtils.NewMockDB(t)
	s := store.NewPSQLGameStore(db)
	query := regexp.QuoteMeta("update gme set status=$2, img_url=$3, map_url=$4, plrs=$5, dur=$6, age_rng=$7, nme=$8, descr=$9, addr=$10 where gme_id = $1;")
	mg := testUtils.DB1.Games

	t.Run("update game successfully", func(t *testing.T) {
		mock.ExpectExec(query).WithArgs(mg[0].GameID, mg[0].Status, mg[0].ImgURL, mg[0].MapURL, mg[0].Plrs, mg[0].Dur, mg[0].AgeRange, mg[0].Name, mg[0].Descr, mg[0].Addr).
			WillReturnResult(sqlmock.NewResult(0, 1))

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		ok, err := s.UpdateGame(ctx, mg[0])

		testUtils.AssertError(t, err, false)
		testUtils.AssertEqual(t, ok, true)
	})

	t.Run("update game unsuccessfully", func(t *testing.T) {
		mock.ExpectExec(query).WithArgs(mg[0].GameID, mg[0].Status, mg[0].ImgURL, mg[0].MapURL, mg[0].Plrs, mg[0].Dur, mg[0].AgeRange, mg[0].Name, mg[0].Descr, mg[0].Addr).
			WillReturnResult(sqlmock.NewResult(0, 0))

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		ok, err := s.UpdateGame(ctx, mg[0])

		testUtils.AssertError(t, err, false)
		testUtils.AssertEqual(t, ok, false)
	})

}
