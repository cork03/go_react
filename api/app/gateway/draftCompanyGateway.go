package gateway

import (
	"go-rest-api/domain"
	"go-rest-api/driverBoundary"
	"go-rest-api/gatewayBoundary"
	"go-rest-api/model"
)

type draftCompanyGateway struct {
	DraftCompanyDriver driverBoundary.IDraftCompanyDriver
}

func (draftCompanyGateway *draftCompanyGateway) Create(company domain.Company, mailCertificationId uint) error {
	draftCompanyModel := model.DraftCompany{
		MailCertificationID: mailCertificationId,
		Name:                company.Name,
		PostalCode:          company.PostalCode,
		Prefecture:          company.Prefecture,
		Town:                company.Town,
		Area:                company.Area,
		Tel:                 company.Tel,
	}
	if err := draftCompanyGateway.DraftCompanyDriver.Create(draftCompanyModel); err != nil {
		return err
	}
	return nil
}

func NewDraftCompanyGateway(draftCompanyDriver driverBoundary.IDraftCompanyDriver) gatewayBoundary.IDraftCompanyGateway {
	return &draftCompanyGateway{draftCompanyDriver}
}
