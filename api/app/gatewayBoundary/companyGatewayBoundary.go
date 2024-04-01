package gatewayBoundary

import "go-rest-api/domain"

type ICompanyGateway interface {
	Register(company domain.Company) error
}
