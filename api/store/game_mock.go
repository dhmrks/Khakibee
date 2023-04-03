package store

import (
	"context"

	"gitlab.com/khakibee/khakibee/api/model"
)

type MockGameStore struct {
	Games        []model.Game
	Calendar     map[int][]model.Event
	InsertCalled int
	UpdateCalled int
	ReturnError  error
}

func (s *MockGameStore) SelectGames(ctx context.Context, status string) (games []model.Game, err error) {
	if s.ReturnError != nil {
		return nil, s.ReturnError
	}

	if status == GameStatusAll {
		games = s.Games
	} else {
		for _, g := range s.Games {
			if g.Status == status {
				games = append(games, g)
			}
		}
	}

	return
}

func (s *MockGameStore) SelectGameByID(ctx context.Context, gameID int) (game *model.Game, err error) {
	if s.ReturnError != nil {
		return nil, s.ReturnError
	}

	for idx, g := range s.Games {
		if g.GameID == gameID {
			game = &s.Games[idx]
		}
	}

	return
}

func (s *MockGameStore) InsertGame(ctx context.Context, game model.Game) (err error) {
	s.InsertCalled += 1
	if s.ReturnError != nil {
		return s.ReturnError
	}
	return nil
}

func (s *MockGameStore) UpdateGame(ctx context.Context, game model.Game) (ok bool, err error) {
	s.UpdateCalled += 1
	err = s.ReturnError
	if err != nil {
		return
	}

	for _, g := range s.Games {
		if g.GameID == game.GameID {
			return true, nil
		}
	}

	return
}
