package email

type MockEmailHandler struct {
	SendEmailCalledWith SendEmailParams
	ReturnError         error
}

type SendEmailParams struct {
	BkngID    int
	TmpltCode string
	Timezone  string
}

//NewEmailHandler
func NewMockEmailHandler() *MockEmailHandler {
	return &MockEmailHandler{}
}

func (e *MockEmailHandler) SendEmail(bkngID int, tmpltCode string, timezone string) (err error) {
	e.SendEmailCalledWith = SendEmailParams{bkngID, tmpltCode, timezone}
	err = e.ReturnError
	return
}
