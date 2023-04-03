package model

type Game struct {
	GameID int    `json:"game_id"`
	Status string `json:"status"`

	Name  string `json:"name"`
	Descr string `json:"descr"`
	Addr  string `json:"addr"`

	ImgURL   string `json:"img_url"`
	MapURL   string `json:"map_url"`
	Plrs     string `json:"players"`
	Dur      int    `json:"duration"`
	AgeRange string `json:"age_range"`
}
