package mail

import "time"

type IMailSendDriver interface {
	SendMailCertification(email string, token string, expire time.Time) error
}
