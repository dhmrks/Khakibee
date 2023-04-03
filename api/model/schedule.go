package model

type Schedule struct {
	SchdlID  int          `json:"-"`
	WkSchdl  WeekSchedule `json:"week_schedule"`
	Dur      int          `json:"dur"`
	ActiveBy string       `json:"active_by"`
}

type WeekSchedule struct {
	Sunday    []string `json:"0,omitempty"`
	Monday    []string `json:"1,omitempty"`
	Tuesday   []string `json:"2,omitempty"`
	Wednesday []string `json:"3,omitempty"`
	Thursday  []string `json:"4,omitempty"`
	Friday    []string `json:"5,omitempty"`
	Saturday  []string `json:"6,omitempty"`
}

type ScheduleResp struct {
	WeekSchedule
	Upcoming *UpcomingSchdl `json:"upcoming,omitempty"`
}

type UpcomingSchdl struct {
	WeekSchedule
	ActiveBy string `json:"active_by"`
}
