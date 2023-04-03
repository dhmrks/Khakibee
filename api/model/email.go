package model

// EmailTmpl  model
type EmailTmpl struct {
	Subject string
	Body    string
}

// EmailDtls model
type EmailDtls struct {
	GameName string
	Name     string
	Email    string
	MobNum   string
	Date     string
	Notes    string
}

type SMTPDtls struct {
	Username string
	Password string
	Host     string
	Port     string
}
