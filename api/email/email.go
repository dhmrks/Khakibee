package email

import (
	"bytes"
	"context"
	"fmt"
	"net/smtp"
	"text/template"
	"time"

	"gitlab.com/khakibee/khakibee/api/config"
	"gitlab.com/khakibee/khakibee/api/model"
	"gitlab.com/khakibee/khakibee/api/store"
)

type EmailHandler interface {
	SendEmail(bkngID int, tmpltCode string, timezone string) (err error)
}

const (
	TempleteNewBkng    = "newbkng"
	TempleteUpdateBkng = "updbkng"
	TempleteCancelBkng = "cncbkng"
)

type SMTPEmailHandler struct {
	emailStore  store.EmailStore
	smtpDtls    config.SMTPDtls
	authEnabled bool
}

//NewSMTPEmailHandler
func NewSMTPEmailHandler(emailStore store.EmailStore, smtpDtls config.SMTPDtls) (*SMTPEmailHandler, error) {
	email, err := generateEmail(smtpDtls.Username, model.EmailTmpl{Body: "Verifing SMTP configuration", Subject: "Verifing SMTP configuration"}, model.EmailDtls{})
	if err != nil {
		return nil, err
	}

	err = send([]string{smtpDtls.Username}, email, smtpDtls, true)
	return &SMTPEmailHandler{emailStore, smtpDtls, true}, err
}

//NewSMTPEmailHandlerTest for testing purposes. Auth is not supported
func NewSMTPEmailHandlerTest(emailStore store.EmailStore, smtpDtls config.SMTPDtls) *SMTPEmailHandler {
	return &SMTPEmailHandler{emailStore, smtpDtls, false}
}

func (e *SMTPEmailHandler) SendEmail(bkngID int, tmpltCode string, timezone string) (err error) {
	ctc := context.Background()
	ctx, cancel := context.WithTimeout(ctc, 2*time.Second)
	defer cancel()

	dtls, err := e.emailStore.SelectEmailDtls(ctx, bkngID, timezone)
	if err != nil {
		return
	}

	tmplt, err := e.emailStore.SelectEmailTemplate(ctx, tmpltCode)
	if err != nil {
		return
	}

	emailMsg, err := generateEmail(e.smtpDtls.Username, tmplt, dtls)
	if err != nil {
		return
	}

	to := []string{dtls.Email}
	err = send(to, emailMsg, e.smtpDtls, e.authEnabled)

	e.emailStore.InsertEmailReport(ctx, bkngID, tmpltCode, err)

	return
}

func send(to []string, emailMsg []byte, smtpDtls config.SMTPDtls, authEnabled bool) (err error) {
	var user = smtpDtls.Username
	var pass = smtpDtls.Password
	var host = smtpDtls.Host
	var addr = host + ":" + smtpDtls.Port

	var auth smtp.Auth
	if authEnabled {
		auth = smtp.PlainAuth("", user, pass, host)
	}

	err = smtp.SendMail(addr, auth, user, to, emailMsg)

	return
}

// generateEmail parse and execute HTML templates safely
func generateEmail(from string, tmplt model.EmailTmpl, dtls model.EmailDtls) (email []byte, err error) {
	const mime = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"
	message := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\n%s\n\n\n%s",
		from, dtls.Email, tmplt.Subject, mime, tmplt.Body)

	t := template.New("html")
	t, err = t.Parse(message)
	if err != nil {
		err = fmt.Errorf("parse emails template error: %s", err)
		return nil, err
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, dtls)
	email = buf.Bytes()
	return
}
