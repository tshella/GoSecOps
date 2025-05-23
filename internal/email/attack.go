package email

import (
	"fmt"
	"net/smtp"
)

type EmailAttackRequest struct {
	From    string
	To      string
	Subject string
	Body    string
}

func SendSpoofedEmail(req EmailAttackRequest) error {
	msg := []byte(fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s",
		req.From, req.To, req.Subject, req.Body))

	auth := smtp.PlainAuth("", "", "", "localhost") // Use MailDev or Mailhog
	return smtp.SendMail("localhost:1025", auth, req.From, []string{req.To}, msg)
}
