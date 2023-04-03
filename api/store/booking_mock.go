package store

import (
	"context"

	"gitlab.com/khakibee/khakibee/api/model"
)

type MockBookingStore struct {
	Bookings           []model.Booking
	InsertCalled       int
	InsertCalledWithID int
	UpdateCalled       int
	UpdateCalledWithID int
	CheckCalled        int
	CancelCalled       int
	ReturnError        error
	NewBookingID       int
}

func (s *MockBookingStore) SelectBookings(ctx context.Context, gme int, timezone string, offset int) (bkngs []model.Booking, err error) {
	err = s.ReturnError
	if err != nil {
		return
	}

	// now := time.Now()
	// currentYear, currentMonth, _ := now.Date()
	// currentLocation := now.Location()

	// firstOfMonthTm := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	// firstOfMonth := firstOfMonthTm.Format("2006-01-02 15:04")
	// lastOfMonth := firstOfMonthTm.AddDate(0, 1, -1).Format("2006-01-02 15:04")

	for idx, b := range s.Bookings {
		if b.GameID == gme {
			bkngs = append(bkngs, s.Bookings[idx])
		}
	}
	return
}

func (s *MockBookingStore) SelectBookingByID(ctx context.Context, gmeID int, bkngID int) (bkng *model.Booking, err error) {
	err = s.ReturnError
	if err != nil {
		return
	}

	for idx, b := range s.Bookings {
		if b.GameID == gmeID && b.BkngID == bkngID {
			bkng = &s.Bookings[idx]
		}
	}

	return
}

func (s *MockBookingStore) UpdateBooking(ctx context.Context, bkng model.Booking, lstEditedBy int) (notFound, dateConflict bool, err error) {
	s.UpdateCalled += 1
	err = s.ReturnError
	s.UpdateCalledWithID = lstEditedBy
	notFound = true
	if err != nil {
		return
	}

	for _, b := range s.Bookings {
		if b.BkngID != bkng.BkngID && b.Dte == bkng.Dte && b.GameID == bkng.GameID && b.Status == "booked" {
			dateConflict = true
		}
		if b.BkngID == bkng.BkngID {
			notFound = false
		}
	}

	return
}

func (s *MockBookingStore) CancelBooking(ctx context.Context, gmeID int, bkngID int) (ok bool, err error) {
	s.CancelCalled += 1
	err = s.ReturnError
	if err != nil {
		return
	}

	for _, b := range s.Bookings {
		if b.GameID == gmeID && b.BkngID == bkngID {
			ok = true
		}
	}

	return
}

func (s *MockBookingStore) InsertBooking(ctx context.Context, bkng model.Booking, bkedBy int) (bkngID *int, err error) {
	s.InsertCalled += 1
	err = s.ReturnError
	s.InsertCalledWithID = bkedBy
	var exist bool
	if err != nil {
		return
	}

	for _, b := range s.Bookings {
		if b.Dte == bkng.Dte && b.GameID == bkng.GameID && b.Status == "booked" {
			exist = true
		}
	}

	if !exist {
		bkngID = new(int)
		*bkngID = s.NewBookingID
		s.Bookings = append(s.Bookings, bkng)
	}

	return
}
