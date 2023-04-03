package model

type Booking struct {
	BkngID int    `json:"bkng_id"`
	Status string `json:"status,omitempty"`

	Dte    string `json:"date,omitempty"`
	GameID int    `json:"game_id,omitempty"`

	Name   string `json:"name"`
	MobNum string `json:"mob_num,omitempty"`
	Email  string `json:"email,omitempty"`

	Notes *string `json:"notes,omitempty"`
}
