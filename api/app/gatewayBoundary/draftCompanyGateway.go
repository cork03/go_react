package gatewayBoundary

import "go-rest-api/domain"

type IDraftCompanyGateway interface {
	Create(company domain.Company, mailCertificationId uint) error
}
