package mail

import (
	mailDriver "go-rest-api/driverBoundary/mail"
	"go-rest-api/gatewayBoundary/mail"
	"time"
)

type MailSendGateway struct {
	mailSendDriver mailDriver.IMailSendDriver
}

func (sendGateway MailSendGateway) SendMailCertification(email string, token string, expire time.Time) error {
	if err := sendGateway.mailSendDriver.SendMailCertification(email, token, expire); err != nil {
		return err
	}
	return nil
}

func NewMailSendGateway(mailSendDriver mailDriver.IMailSendDriver) mail.IMailSendGateway {
	return &MailSendGateway{mailSendDriver: mailSendDriver}
}
