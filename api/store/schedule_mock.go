package store

import (
	"context"

	"gitlab.com/khakibee/khakibee/api/model"
)

type MockScheduleStore struct {
	Calendar map[int][]model.Event
	Schedule map[int][]*model.Schedule

	CheckCalled                int
	SelectCalled               int
	UpsertCalledWithSchdlID    *int
	UpsertNewCalledWithSchdlID *int
	ReturnError                error
}

func (s *MockScheduleStore) UpsertSchedule(ctx context.Context, gmeID int, schdl model.WeekSchedule, schdlID *int) (err error) {
	s.UpsertCalledWithSchdlID = schdlID
	err = s.ReturnError
	return
}

func (s *MockScheduleStore) UpsertNewSchedule(ctx context.Context, gmeID int, schdl model.WeekSchedule, activeBy string, schdlID *int) (err error) {
	s.UpsertNewCalledWithSchdlID = schdlID
	err = s.ReturnError
	return
}

func (s *MockScheduleStore) SelectScheduleByGameID(ctx context.Context, gmID int, upcoming bool) (schdl *model.Schedule, err error) {
	s.SelectCalled += 1
	err = s.ReturnError
	if err != nil {
		return
	}

	if s.Schedule[gmID] != nil {
		if upcoming && len(s.Schedule[gmID]) > 1 {
			schdl = s.Schedule[gmID][1]
		} else if !upcoming {
			schdl = s.Schedule[gmID][0]
		}
	}

	return
}

func (s *MockScheduleStore) CheckSchedule(ctx context.Context, date string, gameID int) (exists bool, err error) {
	s.CheckCalled += 1
	if s.ReturnError != nil {
		return false, s.ReturnError
	}

	for _, e := range s.Calendar[gameID] {
		if e.Start == date {
			exists = true
			return
		}
	}

	return
}
