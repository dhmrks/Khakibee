package email_test

import (
	"errors"
	"fmt"
	"strconv"
	"testing"

	"gitlab.com/khakibee/khakibee/api/config"
	"gitlab.com/khakibee/khakibee/api/email"
	"gitlab.com/khakibee/khakibee/api/store"
	"gitlab.com/khakibee/khakibee/api/testUtils"

	smtpmock "github.com/mocktools/go-smtp-mock"
)

var (
	es       *store.MockEmailStore
	h        *email.SMTPEmailHandler
	errStore = errors.New("internal server error")
)

var server = smtpmock.New(smtpmock.ConfigurationAttr{
	LogToStdout:       false,
	LogServerActivity: false,
})

const EmailMsg = "From: ioulios.tsiko@gmail.com\r\nTo: ioulios@email.com\r\nSubject: H κράτηση στο Academy\r\nMIME-version: 1.0;\r\nContent-Type: text/html; charset=\"UTF-8\";\r\n\r\n\r\n<!DOCTYPE html PUBLIC \"-//W3C//DTD HTML 4.01 Transitional//EN\" \"http://www.w3.org/TR/html4/loose.dtd\"><html><body><h3>Αγαπητέ/τή&nbsp;Ioulios Tsiko,</h3><p>Σας ευχαριστούμε για την κράτηση σας στο Paradox Project:<b>Η Έπαυλη</b>! Έχετε κάνει κράτηση στις<b>2022-07-28 14:00</b>.</p><p>Καλή διασκέδαση!</p><p>Ευχαριστούμε.</p></body>undefined</html>\r\n"

func setup() {

	// To start server use Start() method
	if err := server.Start(); err != nil {
		fmt.Println(err)
	}

	hostAddress, portNumber := "127.0.0.1", server.PortNumber

	var smtpDtls = config.SMTPDtls{
		Username: "ioulios.tsiko@gmail.com",
		Password: "zpwswxcpjdedrfjr",
		Host:     hostAddress,
		Port:     strconv.Itoa(portNumber),
	}

	es = &store.MockEmailStore{EmailDtls: testUtils.DB1.EmailDtls, EmailTmpl: testUtils.DB1.EmailTmpl}
	h = email.NewSMTPEmailHandlerTest(es, smtpDtls)
}

func TestSendEmail(t *testing.T) {
	setup()

	t.Run("send email successfully", func(t *testing.T) {
		err := h.SendEmail(1, "correct", "athens")

		testUtils.AssertError(t, err, false)
		testUtils.AssertFunctionCalled(t, es.SelectEmailDtlsCalled, 1, "SelectEmailDtls")
		testUtils.AssertFunctionCalled(t, es.SelectEmailTmplCalled, 1, "SelectEmailTemplate")
		testUtils.AssertFunctionCalled(t, es.InsertEmailReportCalled, 1, "InsertEmailReport")

		s := server.Messages()
		gotEmailMsg := s[len(s)-1].MsgRequest()

		testUtils.AssertEqual(t, gotEmailMsg, EmailMsg)
	})

	t.Run("send email with sql error", func(t *testing.T) {

		es.ReturnError = errStore
		err := h.SendEmail(1, "correct", "athens")

		testUtils.AssertError(t, err, true)
		testUtils.AssertFunctionCalled(t, es.SelectEmailDtlsCalled, 2, "SelectEmailDtls")
	})

	t.Run("send email with email template error", func(t *testing.T) {
		es.ReturnError = nil
		err := h.SendEmail(1, "errored", "athens")

		testUtils.AssertError(t, err, true)
		testUtils.AssertFunctionCalled(t, es.SelectEmailDtlsCalled, 3, "SelectEmailDtls")
		testUtils.AssertFunctionCalled(t, es.SelectEmailTmplCalled, 2, "SelectEmailTemplate")
	})

	t.Run("send email with smtp error", func(t *testing.T) {
		h = email.NewSMTPEmailHandlerTest(es, config.SMTPDtls{})

		err := h.SendEmail(1, "correct", "athens")

		testUtils.AssertError(t, err, true)
		testUtils.AssertFunctionCalled(t, es.SelectEmailDtlsCalled, 4, "SelectEmailDtls")
		testUtils.AssertFunctionCalled(t, es.SelectEmailTmplCalled, 3, "SelectEmailTemplate")
		testUtils.AssertFunctionCalled(t, es.InsertEmailReportCalled, 2, "InsertEmailReport")
	})

}
