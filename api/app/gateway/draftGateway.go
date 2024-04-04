package gateway

import (
	"go-rest-api/domain"
	"go-rest-api/driverBoundary"
	"go-rest-api/gatewayBoundary"
	"go-rest-api/model"
)

type draftGateway struct {
	DraftCompanyDriver driverBoundary.IDraftDriver
}

func (draftCompanyGateway *draftGateway) Create(
	company domain.Company,
	user domain.User,
	userPassword domain.UserPassword,
	mailCertificationId uint,
) (domain.User, error) {
	draftCompanyModel := model.DraftCompany{
		MailCertificationID: mailCertificationId,
		Name:                company.Name,
		PostalCode:          company.PostalCode,
		Prefecture:          company.Prefecture,
		Town:                company.Town,
		Area:                company.Area,
		Tel:                 company.Tel,
	}
	draftUserModel := model.DraftUser{
		MailCertificationID: mailCertificationId,
		Email:               user.Email,
		Name:                user.Name,
	}
	draftUserPassword := model.DraftUserPassword{
		MailCertificationID: mailCertificationId,
		Password:            userPassword.GetPassword(),
	}
	resDraftUser, err := draftCompanyGateway.DraftCompanyDriver.Create(
		draftCompanyModel,
		draftUserModel,
		draftUserPassword,
	)
	if err != nil {
		return domain.User{}, err
	}
	resUser := domain.User{
		Email: resDraftUser.Email,
		Name:  resDraftUser.Name,
	}
	return resUser, nil
}

func NewDraftGateway(draftCompanyDriver driverBoundary.IDraftDriver) gatewayBoundary.IDraftGateway {
	return &draftGateway{draftCompanyDriver}
}
