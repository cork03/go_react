package mail

import (
	"fmt"
	"go-rest-api/driverBoundary/mail"
	"net/smtp"
	"strings"
)

type MailSendDriver struct {
}

var (
	hostname = "localhost"
	port     = 1025
)

func (m MailSendDriver) SendMailCertification(email string) error {
	from := "sender@example.com"
	subject := "Hello"
	body := "Hello\nURL"

	msg := []byte(strings.ReplaceAll(fmt.Sprintf("To: %s\nSubject: %s\n\n%s", email, subject, body), "\n", "\r\n"))
	if err := smtp.SendMail(fmt.Sprintf("%s:%d", hostname, port), nil, from, []string{email}, msg); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func NewMailSendDriver() mail.IMailSendDriver {
	return &MailSendDriver{}
}
