package store

import (
	"context"
	"database/sql"

	"github.com/neko-neko/echo-logrus/v2/log"
	"github.com/sirupsen/logrus"
	"gitlab.com/khakibee/khakibee/api/model"
)

type GameStore interface {
	SelectGames(ctx context.Context, status string) (games []model.Game, err error)
	SelectGameByID(ctx context.Context, gameID int) (game *model.Game, err error)
	InsertGame(ctx context.Context, game model.Game) (err error)
	UpdateGame(ctx context.Context, game model.Game) (ok bool, err error)
}

type PSQLGameStore struct {
	db *sql.DB
}

func NewPSQLGameStore(db *sql.DB) (gameStore *PSQLGameStore) {
	gameStore = &PSQLGameStore{db}

	return
}

const (
	GameStatusAll      = "all"
	GameStatusActive   = "active"
	GameStatusInactive = "inactive"
)

func (g *PSQLGameStore) SelectGames(ctx context.Context, status string) (games []model.Game, err error) {
	query := `select
		g.gme_id, g.status, g.img_url, g.map_url, g.plrs, g.dur, g.age_rng,
		g.nme, g.descr, g.addr
	from gme g`

	var args []interface{}
	if status != GameStatusAll {
		query += `
			where status = $1`
		args = append(args, status)
	}

	rows, err := g.db.QueryContext(ctx, query, args...)
	if err != nil {
		log.Logger().WithFields(logrus.Fields{"query": query}).Error(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var game model.Game
		err = rows.Scan(
			&game.GameID,
			&game.Status,
			&game.ImgURL,
			&game.MapURL,
			&game.Plrs,
			&game.Dur,
			&game.AgeRange,
			&game.Name,
			&game.Descr,
			&game.Addr,
		)

		if err != nil {
			return nil, err
		}
		games = append(games, game)
	}

	return
}

func (g *PSQLGameStore) SelectGameByID(ctx context.Context, gameID int) (game *model.Game, err error) {
	query := `select 
		g.gme_id, g.status, g.img_url, g.map_url, g.plrs, g.dur, g.age_rng,
		g.nme, g.descr, g.addr
	from gme g
	where g.gme_id = $1;`

	args := []interface{}{gameID}

	row := g.db.QueryRowContext(ctx, query, args...)
	var gm model.Game
	err = row.Scan(
		&gm.GameID,
		&gm.Status,
		&gm.ImgURL,
		&gm.MapURL,
		&gm.Plrs,
		&gm.Dur,
		&gm.AgeRange,
		&gm.Name,
		&gm.Descr,
		&gm.Addr,
	)

	if err == sql.ErrNoRows {
		err = nil
	} else if err != nil {
		log.Logger().WithFields(logrus.Fields{"query": query, "args": args}).Error(err)
	} else {
		game = &gm
	}

	return
}

func (g *PSQLGameStore) InsertGame(ctx context.Context, game model.Game) (err error) {
	query := `insert into gme (status, img_url, map_url, plrs, dur, age_rng, nme, descr, addr) 
	values ($1, $2, $3, $4, $5, $6, $7, $8, $9);`

	args := []interface{}{game.Status, game.ImgURL, game.MapURL, game.Plrs, game.Dur, game.AgeRange, game.Name, game.Descr, game.Addr}

	_, err = g.db.ExecContext(ctx, query, args...)
	if err != nil {
		log.Logger().WithFields(logrus.Fields{"query": query, "args": args}).Error(err)
	}

	return
}

func (g *PSQLGameStore) UpdateGame(ctx context.Context, game model.Game) (ok bool, err error) {
	query := `update gme  
	set status=$2, img_url=$3, map_url=$4, plrs=$5, dur=$6, age_rng=$7, nme=$8, descr=$9, addr=$10
	where gme_id = $1;`

	args := []interface{}{game.GameID, game.Status, game.ImgURL, game.MapURL, game.Plrs, game.Dur, game.AgeRange, game.Name, game.Descr, game.Addr}

	res, err := g.db.ExecContext(ctx, query, args...)
	if err != nil {
		log.Logger().WithFields(logrus.Fields{"query": query, "args": args}).Error(err)
		return false, err
	}

	ra, err := res.RowsAffected()
	ok = ra == 1
	return
}
