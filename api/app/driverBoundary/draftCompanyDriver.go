package driverBoundary

import "go-rest-api/model"

type IDraftCompanyDriver interface {
	Create(company model.DraftCompany) error
}
