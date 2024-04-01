package gateway

import (
	"go-rest-api/domain"
	"go-rest-api/driverBoundary"
	"go-rest-api/gatewayBoundary"
	"go-rest-api/model"
)

type CompanyGateway struct {
	companyDriver driverBoundary.ICompanyDriver
}

func (companyGateway *CompanyGateway) Register(company domain.Company) error {
	companyModel := model.Company{
		Name:       company.Name,
		PostalCode: company.PostalCode,
		Prefecture: company.Prefecture,
		Town:       company.Town,
		Area:       company.Area,
		Tel:        company.Tel,
	}
	if err := companyGateway.companyDriver.Register(companyModel); err != nil {
		return err
	}
	return nil
}

func NewCompanyGateway(companyDriver driverBoundary.ICompanyDriver) gatewayBoundary.ICompanyGateway {
	return &CompanyGateway{companyDriver: companyDriver}
}
