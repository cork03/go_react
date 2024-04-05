package mail

import (
	"fmt"
	"go-rest-api/driverBoundary/mail"
	"net/smtp"
	"os"
	"strings"
)

type MailSendDriver struct {
}

var (
	hostname = "localhost"
	port     = 1025
)

func (m MailSendDriver) SendMailCertification(email string, token string) error {
	from := os.Getenv("MAIL_SENDER")
	subject := "メール認証のお知らせ"
	body := fmt.Sprintf("下記URLよりメールの認証をお願いいたします。\nhttp://%s/mail-certification?token=%s", os.Getenv("HOST"), token)

	msg := []byte(strings.ReplaceAll(fmt.Sprintf("To: %s\nSubject: %s\n\n%s", email, subject, body), "\n", "\r\n"))
	if err := smtp.SendMail(fmt.Sprintf("%s:%d", hostname, port), nil, from, []string{email}, msg); err != nil {
		return err
	}
	return nil
}

func NewMailSendDriver() mail.IMailSendDriver {
	return &MailSendDriver{}
}
