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

func (mailCertificationGateway *MailCertificationGateway) GetDraftsByToken(token string) (*domain.Drafts, error) {
	resMailCertificationModel, err := mailCertificationGateway.MailCertificationDriver.GetByTokenWithDrafts(token)
	if err != nil {
		return nil, err
	}
	if resMailCertificationModel == nil {
		return nil, nil
	}

	drafts := domain.Drafts{
		MailCertification: domain.MailCertification{
			ID:     resMailCertificationModel.ID,
			Token:  resMailCertificationModel.Token,
			Expire: resMailCertificationModel.Expire,
		},
		DraftUser: domain.User{
			Name:  resMailCertificationModel.DraftUser.Name,
			Email: resMailCertificationModel.DraftUser.Email,
		},
		DraftCompany: domain.Company{
			Name:       resMailCertificationModel.DraftCompany.Name,
			PostalCode: resMailCertificationModel.DraftCompany.PostalCode,
			Prefecture: resMailCertificationModel.DraftCompany.Prefecture,
			Town:       resMailCertificationModel.DraftCompany.Town,
			Area:       resMailCertificationModel.DraftCompany.Area,
			Tel:        resMailCertificationModel.DraftCompany.Tel,
		},
		DraftUserPassword: domain.DraftUserPassword{
			Password: resMailCertificationModel.DraftUserPassword.Password,
		},
	}

	return &drafts, nil
}

func NewMailCertificationGateway(mailCertificationDriver driverBoundary.IMailCertificationDriver) gatewayBoundary.IMailCertificationGateway {
	return &MailCertificationGateway{mailCertificationDriver}
}
