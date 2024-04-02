package gateway

import (
	"go-rest-api/domain"
	"go-rest-api/driverBoundary"
	"go-rest-api/gatewayBoundary"
	"go-rest-api/model"
)

type MailCertificationGateway struct {
	MailCertificationDriver driverBoundary.IMailCertificationDriver
}

func (mailCertificationGateway *MailCertificationGateway) Create(mailCertification domain.MailCertification) error {
	mailCertificationModel := model.MailCertification{
		Token:  mailCertification.Token,
		Expire: mailCertification.Expire,
	}
	if err := mailCertificationGateway.MailCertificationDriver.Create(mailCertificationModel); err != nil {
		return err
	}
	return nil
}

func NewMailCertificationGateway(mailCertificationDriver driverBoundary.IMailCertificationDriver) gatewayBoundary.IMailCertificationGateway {
	return &MailCertificationGateway{mailCertificationDriver}
}
