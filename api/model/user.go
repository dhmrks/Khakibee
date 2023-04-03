package model

type Auth struct {
	Token string `json:"token"`
}

type User struct {
	UserID int
	Usrnme string
	Psswrd string
}
