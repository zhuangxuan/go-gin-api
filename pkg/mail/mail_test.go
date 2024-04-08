package mail

import (
	"testing"
)

func TestSend(t *testing.T) {
	options := &Options{
		MailHost: "smtp.163.com",
		MailPort: 465,
		MailUser: "17602196562@163.com",
		MailPass: "JPIXSNOGYCZMXWGV", //密码或授权码
		MailTo:   "17602196562@163.com",
		Subject:  "subject",
		Body:     "body",
	}
	err := Send(options)
	if err != nil {
		t.Error("Mail Send error", err)
		return
	}
	t.Log("success")
}
