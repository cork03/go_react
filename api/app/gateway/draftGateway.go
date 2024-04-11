package gateway

import (
	"go-rest-api/domain"
	"go-rest-api/driverBoundary"
	"go-rest-api/gateway/dto"
	"go-rest-api/gatewayBoundary"
	"go-rest-api/model"
)

type draftGateway struct {
	DraftCompanyDriver driverBoundary.IDraftDriver
	CompanyDriver      driverBoundary.ICompanyDriver
}

func (draftCompanyGateway *draftGateway) BookRegistration(drafts domain.Drafts) error {
	draftDto := dto.Drafts{
		User: model.User{
			Email: drafts.DraftUser.Email,
			Name:  drafts.DraftUser.Name,
		},
		Company: model.Company{
			Name:       drafts.DraftCompany.Name,
			PostalCode: drafts.DraftCompany.PostalCode,
			Prefecture: drafts.DraftCompany.Prefecture,
			Town:       drafts.DraftCompany.Town,
			Area:       drafts.DraftCompany.Area,
		},
		UserPassword: model.UserPassword{
			Password: drafts.DraftUserPassword.Password,
		},
	}
	if err := draftCompanyGateway.CompanyDriver.BookRegistration(draftDto); err != nil {
		return err
	}
	return nil
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

func NewDraftGateway(
	draftCompanyDriver driverBoundary.IDraftDriver,
	companyDriver driverBoundary.ICompanyDriver,
) gatewayBoundary.IDraftGateway {
	return &draftGateway{
		draftCompanyDriver,
		companyDriver,
	}
}
