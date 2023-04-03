package testUtils

import (
	"gitlab.com/khakibee/khakibee/api/model"
)

type DB struct {
	Games     []model.Game
	Bookings  []model.Booking
	Users     []model.User
	Schedule  map[int][]*model.Schedule
	Calendar  map[int][]model.Event
	EmailTmpl map[string]model.EmailTmpl
	EmailDtls model.EmailDtls
}

var DB1 = DB{
	Games: []model.Game{
		{
			GameID:   1,
			Status:   "active",
			Name:     "The Mansion",
			Descr:    "A mansion of mystery",
			Addr:     "21 Athens str, Athens, Greece 21012",
			ImgURL:   "booking.com/musicacademy.jpg",
			MapURL:   "https://goo.gl/maps/ghrrS37V7tPGw5Zd8",
			Plrs:     "2-5",
			Dur:      150,
			AgeRange: "13+",
		}, {
			GameID:   2,
			Status:   "active",
			Name:     "Bookstore",
			Descr:    "A scary bookstore",
			Addr:     "21 Athens str, Athens, Greece 21012",
			ImgURL:   "booking.com/musicacademy.jpg",
			MapURL:   "https://goo.gl/maps/ghrrS37V7tPGw5Zd8",
			Plrs:     "2-5",
			Dur:      150,
			AgeRange: "13+",
		}, {
			GameID:   3,
			Status:   "inactive",
			Name:     "Academy",
			Descr:    "A scary bookstore",
			Addr:     "21 Athens str, Athens, Greece 21012",
			ImgURL:   "booking.com/musicacademy.jpg",
			MapURL:   "https://goo.gl/maps/ghrrS37V7tPGw5Zd8",
			Plrs:     "2-5",
			Dur:      150,
			AgeRange: "13+",
		},
	},
	Bookings: []model.Booking{
		{
			BkngID: 1,
			Dte:    "2022-06-30 14:00",
			Name:   "ioulios",
			MobNum: "+306975645865",
			Email:  "ioulios@email.com",
			Status: "booked",
			GameID: 1,
		}, {
			BkngID: 2,
			Dte:    "2022-06-30 19:00",
			Name:   "ioulios",
			MobNum: "+306975645865",
			Email:  "ioulios@email.com",
			Status: "booked",
			GameID: 1,
		}, {
			BkngID: 3,
			Dte:    "2022-06-30 19:00",
			Name:   "ioulios",
			MobNum: "+306975645865",
			Email:  "ioulios@email.com",
			Status: "booked",
			GameID: 2,
		},
	},
	Users: []model.User{
		{
			UserID: 1,
			Usrnme: "ioulios",
			Psswrd: "pass",
		},
	},
	Schedule: map[int][]*model.Schedule{
		1: {
			{
				SchdlID: 1,
				WkSchdl: model.WeekSchedule{
					Sunday:    []string{"14:00", "19:00"},
					Monday:    []string{"14:00", "19:00"},
					Tuesday:   []string{"14:00", "19:00"},
					Wednesday: []string{"14:00", "19:00"},
					Thursday:  []string{"14:00", "19:00"},
					Friday:    []string{"14:00", "19:00"},
					Saturday:  []string{"14:00", "19:00"},
				},
				Dur:      160,
				ActiveBy: "2022-09-20",
			},
			{
				SchdlID: 2,
				WkSchdl: model.WeekSchedule{
					Sunday:    []string{"13:00", "18:00"},
					Monday:    []string{"13:00", "18:00"},
					Tuesday:   []string{"13:00", "18:00"},
					Wednesday: []string{"13:00", "18:00"},
					Thursday:  []string{"13:00", "18:00"},
					Friday:    []string{"13:00", "18:00"},
					Saturday:  []string{"13:00", "18:00"},
				},
				Dur:      160,
				ActiveBy: "2022-10-30",
			},
		},
		2: {
			{
				SchdlID: 3,
				WkSchdl: model.WeekSchedule{
					Sunday:    []string{"14:00", "19:00"},
					Monday:    []string{"14:00", "19:00"},
					Tuesday:   []string{"14:00", "19:00"},
					Wednesday: []string{"14:00", "19:00"},
					Thursday:  []string{"14:00", "19:00"},
					Friday:    []string{"14:00", "19:00"},
					Saturday:  []string{"14:00", "19:00"},
				},
				Dur:      160,
				ActiveBy: "2022-09-20",
			},
		},
	},
	Calendar: map[int][]model.Event{
		1: {
			{Start: "2022-06-01 14:00", End: "2022-06-01 16:40", Booking: nil},
			{Start: "2022-06-01 19:00", End: "2022-06-01 21:40", Booking: nil},
			{Start: "2022-06-02 14:00", End: "2022-06-02 16:40", Booking: nil},
			{Start: "2022-06-02 19:00", End: "2022-06-02 21:40", Booking: nil},
			{Start: "2022-06-03 14:00", End: "2022-06-03 16:40", Booking: nil},
			{Start: "2022-06-03 19:00", End: "2022-06-03 21:40", Booking: nil},
			{Start: "2022-06-04 14:00", End: "2022-06-04 16:40", Booking: nil},
			{Start: "2022-06-04 19:00", End: "2022-06-04 21:40", Booking: nil},
			{Start: "2022-06-05 14:00", End: "2022-06-05 16:40", Booking: nil},
			{Start: "2022-06-05 19:00", End: "2022-06-05 21:40", Booking: nil},
			{Start: "2022-06-06 14:00", End: "2022-06-06 16:40", Booking: nil},
			{Start: "2022-06-06 19:00", End: "2022-06-06 21:40", Booking: nil},
			{Start: "2022-06-07 14:00", End: "2022-06-07 16:40", Booking: nil},
			{Start: "2022-06-07 19:00", End: "2022-06-07 21:40", Booking: nil},
			{Start: "2022-06-08 14:00", End: "2022-06-08 16:40", Booking: nil},
			{Start: "2022-06-08 19:00", End: "2022-06-08 21:40", Booking: nil},
			{Start: "2022-06-09 14:00", End: "2022-06-09 16:40", Booking: nil},
			{Start: "2022-06-09 19:00", End: "2022-06-09 21:40", Booking: nil},
			{Start: "2022-06-10 14:00", End: "2022-06-10 16:40", Booking: nil},
			{Start: "2022-06-10 19:00", End: "2022-06-10 21:40", Booking: nil},
			{Start: "2022-06-11 14:00", End: "2022-06-11 16:40", Booking: nil},
			{Start: "2022-06-11 19:00", End: "2022-06-11 21:40", Booking: nil},
			{Start: "2022-06-12 14:00", End: "2022-06-12 16:40", Booking: nil},
			{Start: "2022-06-12 19:00", End: "2022-06-12 21:40", Booking: nil},
			{Start: "2022-06-13 14:00", End: "2022-06-13 16:40", Booking: nil},
			{Start: "2022-06-13 19:00", End: "2022-06-13 21:40", Booking: nil},
			{Start: "2022-06-14 14:00", End: "2022-06-14 16:40", Booking: nil},
			{Start: "2022-06-14 19:00", End: "2022-06-14 21:40", Booking: nil},
			{Start: "2022-06-15 14:00", End: "2022-06-15 16:40", Booking: nil},
			{Start: "2022-06-15 19:00", End: "2022-06-15 21:40", Booking: nil},
			{Start: "2022-06-16 14:00", End: "2022-06-16 16:40", Booking: nil},
			{Start: "2022-06-16 19:00", End: "2022-06-16 21:40", Booking: nil},
			{Start: "2022-06-17 14:00", End: "2022-06-17 16:40", Booking: nil},
			{Start: "2022-06-17 19:00", End: "2022-06-17 21:40", Booking: nil},
			{Start: "2022-06-18 14:00", End: "2022-06-18 16:40", Booking: nil},
			{Start: "2022-06-18 19:00", End: "2022-06-18 21:40", Booking: nil},
			{Start: "2022-06-19 14:00", End: "2022-06-19 16:40", Booking: nil},
			{Start: "2022-06-19 19:00", End: "2022-06-19 21:40", Booking: nil},
			{Start: "2022-06-20 14:00", End: "2022-06-20 16:40", Booking: nil},
			{Start: "2022-06-20 19:00", End: "2022-06-20 21:40", Booking: nil},
			{Start: "2022-06-21 14:00", End: "2022-06-21 16:40", Booking: nil},
			{Start: "2022-06-21 19:00", End: "2022-06-21 21:40", Booking: nil},
			{Start: "2022-06-22 14:00", End: "2022-06-22 16:40", Booking: nil},
			{Start: "2022-06-22 19:00", End: "2022-06-22 21:40", Booking: nil},
			{Start: "2022-06-23 14:00", End: "2022-06-23 16:40", Booking: nil},
			{Start: "2022-06-23 19:00", End: "2022-06-23 21:40", Booking: nil},
			{Start: "2022-06-24 14:00", End: "2022-06-24 16:40", Booking: nil},
			{Start: "2022-06-24 19:00", End: "2022-06-24 21:40", Booking: nil},
			{Start: "2022-06-25 14:00", End: "2022-06-25 16:40", Booking: nil},
			{Start: "2022-06-25 19:00", End: "2022-06-25 21:40", Booking: nil},
			{Start: "2022-06-26 14:00", End: "2022-06-26 16:40", Booking: nil},
			{Start: "2022-06-26 19:00", End: "2022-06-26 21:40", Booking: nil},
			{Start: "2022-06-27 14:00", End: "2022-06-27 16:40", Booking: nil},
			{Start: "2022-06-27 19:00", End: "2022-06-27 21:40", Booking: nil},
			{Start: "2022-06-28 14:00", End: "2022-06-28 16:40", Booking: nil},
			{Start: "2022-06-28 19:00", End: "2022-06-28 21:40", Booking: nil},
			{Start: "2022-06-29 14:00", End: "2022-06-29 16:40", Booking: nil},
			{Start: "2022-06-29 19:00", End: "2022-06-29 21:40", Booking: nil},
			{Start: "2022-06-30 14:00", End: "2022-06-30 16:40", Booking: &model.Booking{BkngID: 1, Name: "ioulios"}},
			{Start: "2022-06-30 19:00", End: "2022-06-30 21:40", Booking: &model.Booking{BkngID: 2, Name: "ioulios"}},
		},
	},
	EmailTmpl: map[string]model.EmailTmpl{
		"correct": {
			Subject: `H κράτηση στο {{.GameName}}`,
			Body:    `<!DOCTYPE html PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3.org/TR/html4/loose.dtd"><html><body><h3>Αγαπητέ/τή&nbsp;{{.Name}},</h3><p>Σας ευχαριστούμε για την κράτηση σας στο Paradox Project:<b>Η Έπαυλη</b>! Έχετε κάνει κράτηση στις<b>{{.Date}}</b>.</p><p>Καλή διασκέδαση!</p><p>Ευχαριστούμε.</p></body>undefined</html>`,
		},
		"errored": {
			Subject: `H κράτηση στο {{GameName}}`,
			Body:    `<!DOCTYPE html PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3.org/TR/html4/loose.dtd"><html><body><h3>Αγαπητέ/τή&nbsp;{{.Name}},</h3><p>Σας ευχαριστούμε για την κράτηση σας στο Paradox Project:<b>Η Έπαυλη</b>! Έχετε κάνει κράτηση στις<b>{{.Date}}</b>.</p><p>Καλή διασκέδαση!</p><p>Ευχαριστούμε.</p></body>undefined</html>`,
		},
	},
	EmailDtls: model.EmailDtls{
		GameName: "Academy",
		Name:     "Ioulios Tsiko",
		Email:    "ioulios@email.com",
		MobNum:   "+306975645865",
		Date:     "2022-07-28 14:00",
		Notes:    "some note",
	},
}
