package driver

import (
	"go-rest-api/driverBoundary"
	"go-rest-api/model"
	"gorm.io/gorm"
)

type CompanyDriver struct {
	db *gorm.DB
}

func (companyDriver *CompanyDriver) Register(company model.Company) error {
	if err := companyDriver.db.Create(&company).Error; err != nil {
		return err
	}
	return nil
}

func NewCompanyDriver(db *gorm.DB) driverBoundary.ICompanyDriver {
	return &CompanyDriver{db}
}
