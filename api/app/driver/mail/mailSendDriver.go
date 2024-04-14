package mail

import (
	"fmt"
	"go-rest-api/driverBoundary/mail"
	"net/smtp"
	"os"
	"strings"
	"time"
)

type MailSendDriver struct {
}

var (
	hostname = "localhost"
	port     = 1025
)

func (m MailSendDriver) SendMailCertification(email string, token string, expire time.Time) error {
	formattedExpire := expire.Format("2006-01-02 15:04:05")
	from := os.Getenv("MAIL_SENDER")
	subject := "メール認証のお知らせ"
	body := fmt.Sprintf("下記URLよりメールの認証をお願いいたします。\nhttp://%s/mail-certification?token=%s\n本メールの有効期限は%sまでです。", os.Getenv("HOST"), token, formattedExpire)

	msg := []byte(strings.ReplaceAll(fmt.Sprintf("To: %s\nSubject: %s\n\n%s", email, subject, body), "\n", "\r\n"))
	if err := smtp.SendMail(fmt.Sprintf("%s:%d", hostname, port), nil, from, []string{email}, msg); err != nil {
		return err
	}
	return nil
}

func NewMailSendDriver() mail.IMailSendDriver {
	return &MailSendDriver{}
}
