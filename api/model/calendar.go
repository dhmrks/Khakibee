package model

type Event struct {
	Start         string   `json:"start"`
	End           string   `json:"end"`
	Booking       *Booking `json:"bkng,omitempty"`
	OutOfSchedule bool     `json:"out_of_schedule,omitempty"`
}
