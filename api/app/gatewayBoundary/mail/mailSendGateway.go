package mail

import "time"

type IMailSendGateway interface {
	SendMailCertification(email string, token string, expire time.Time) error
}
