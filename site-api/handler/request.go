package handler

// ActiveRoom type
type ActiveRoom struct {
	Room   string `json:"room_id"`
	Active string `json:"active"`
}

type booking struct {
	GameSessID string  `json:"game_sess_id"`
	RoomID     int     `json:"room_id"`
	RoomName   string  `json:"roomName"`
	RoomName2  string  `json:"roomName2"`
	Name       string  `json:"name"`
	MobNum     string  `json:"mob_num"`
	EmailAddr  string  `json:"email_addr"`
	PlrNum     string  `json:"plr_num"`
	PlrLvl     string  `json:"plr_lvl"`
	Age        string  `json:"age"`
	Code       *string `json:"code"`
	LrnedUs    string  `json:"learned_us"`
	TeamName   *string `json:"team_name"`
	PlayRoom1  bool    `json:"playRoom1"`
	EscRoom1   bool    `json:"escRoom1"`
	Diff       string  `json:"diff"`
	Lang       *string `json:"lang"`
	Notes      *string `json:"notes"`
	Date       string  `json:"date"`
	Hour       string  `json:"hour"`
	Status     int     `json:"bkng_status_id"`
	GDPR       bool    `json:"gdpr"`
}

var templateMap = map[int]map[string]map[int]string{
	1: {
		"el": {
			1: "bkng_room1",
			2: "bkng_room2",
			3: "bkng_room3",
		},
		"en": {
			1: "bkng_room1",
			2: "bkng_room2",
			3: "bkng_room3",
		},
	},
	2: {
		"en": {
			1: "edit",
		},
	},
}
