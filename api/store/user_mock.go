package store

import (
	"context"

	"gitlab.com/khakibee/khakibee/api/model"
)

type MockUserStore struct {
	Users       []model.User
	ReturnError error
}

func NewMockUserStore(users []model.User) (userStore *MockUserStore) {
	userStore = &MockUserStore{Users: users}
	return
}

func (s *MockUserStore) SelectUserByCred(ctx context.Context, usrnme, psswrd string) (userID *int, err error) {

	if s.ReturnError != nil {
		return nil, s.ReturnError
	}

	for _, u := range s.Users {
		if u.Usrnme == usrnme && u.Psswrd == psswrd {
			return &u.UserID, nil
		}
	}

	return nil, nil
}
