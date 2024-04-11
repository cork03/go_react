package driverBoundary

import (
	"go-rest-api/gateway/dto"
	"go-rest-api/model"
)

type ICompanyDriver interface {
	Register(company model.Company) error
	BookRegistration(drafts dto.Drafts) error
}
