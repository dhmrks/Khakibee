package store_test

import (
	"context"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"gitlab.com/khakibee/khakibee/api/model"
	"gitlab.com/khakibee/khakibee/api/store"
	"gitlab.com/khakibee/khakibee/api/testUtils"
)

var selectAuthUserRows = []string{"usr_id"}
var mockUsers = []model.User{
	{
		UserID: 1,
		Usrnme: "ioulios",
		Psswrd: "pass",
	},
}

func TestSelectUserByCred(t *testing.T) {
	db, mock := testUtils.NewMockDB(t)
	s := store.NewPSQLUserStore(db)
	query := regexp.QuoteMeta(`select usr_id from usr where usrnme = $1 and psswrd = encode(sha256(($2||salt)::bytea), 'hex');`)
	mu := mockUsers[0]

	t.Run("select user successfully", func(t *testing.T) {
		expectedUID := new(int)
		*expectedUID = 1
		rows := sqlmock.NewRows(selectAuthUserRows).
			AddRow(1)

		mock.ExpectQuery(query).WithArgs(mu.Usrnme, mu.Psswrd).WillReturnRows(rows)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		userID, err := s.SelectUserByCred(ctx, mu.Usrnme, mu.Psswrd)

		testUtils.AssertError(t, err, false)
		testUtils.AssertEqual(t, userID, expectedUID)
	})

	t.Run("select user with no rows", func(t *testing.T) {
		rows := sqlmock.NewRows(selectAuthUserRows)

		mock.ExpectQuery(query).WithArgs(mu.Usrnme, mu.Psswrd).WillReturnRows(rows)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		userID, err := s.SelectUserByCred(ctx, mu.Usrnme, mu.Psswrd)

		testUtils.AssertError(t, err, false)
		testUtils.AssertNil(t, userID)
	})

	t.Run("select user with context deadline exceeded", func(t *testing.T) {
		rows := sqlmock.NewRows(selectAuthUserRows).
			AddRow(nil)

		mock.ExpectQuery(query).WithArgs(mu.Usrnme, mu.Psswrd).WillReturnRows(rows)

		ctx, cancel := context.WithTimeout(context.Background(), 0)
		defer cancel()
		_, err := s.SelectUserByCred(ctx, mu.Usrnme, mu.Psswrd)

		testUtils.AssertError(t, err, true)
	})
}
