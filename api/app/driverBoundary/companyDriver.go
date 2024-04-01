package driverBoundary

import "go-rest-api/model"

type ICompanyDriver interface {
	Register(company model.Company) error
}
