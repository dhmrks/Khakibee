package store

import (
	"context"
	"database/sql"

	"github.com/neko-neko/echo-logrus/v2/log"
	"github.com/sirupsen/logrus"
)

type UserStore interface {
	SelectUserByCred(ctx context.Context, usrnme, psswrd string) (userID *int, err error)
}

type PSQLUserStore struct {
	db *sql.DB
}

func NewPSQLUserStore(db *sql.DB) (userStore *PSQLUserStore) {
	userStore = &PSQLUserStore{db}

	return
}

func (u *PSQLUserStore) SelectUserByCred(ctx context.Context, usrnme, psswrd string) (userID *int, err error) {
	query := `select usr_id  
	from usr 
	where usrnme = $1 
		and psswrd = encode(sha256(($2||salt)::bytea), 'hex');
	`

	args := []interface{}{usrnme, psswrd}

	row := u.db.QueryRowContext(ctx, query, args...)
	err = row.Scan(&userID)
	if err == sql.ErrNoRows {
		args[1] = "*****"
		log.Logger().WithFields(logrus.Fields{"query": query, "args": args}).Error(err)
		err = nil
	}

	return
}
