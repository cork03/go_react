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

func (mailCertificationGateway *MailCertificationGateway) Create(mailCertification domain.MailCertification) (domain.MailCertification, error) {
	mailCertificationModel := model.MailCertification{
		Token:  mailCertification.Token,
		Expire: mailCertification.Expire,
	}
	resMailCertificationModel, err := mailCertificationGateway.MailCertificationDriver.Create(mailCertificationModel)
	if err != nil {
		return domain.MailCertification{}, err
	}
	resMailCertification := domain.MailCertification{
		ID:     resMailCertificationModel.ID,
		Token:  resMailCertificationModel.Token,
		Expire: resMailCertificationModel.Expire,
	}
	return resMailCertification, nil
}

func NewMailCertificationGateway(mailCertificationDriver driverBoundary.IMailCertificationDriver) gatewayBoundary.IMailCertificationGateway {
	return &MailCertificationGateway{mailCertificationDriver}
}
