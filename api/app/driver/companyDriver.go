package driver

import (
	"go-rest-api/driverBoundary"
	"go-rest-api/gateway/dto"
	"go-rest-api/model"
	"gorm.io/gorm"
)

type CompanyDriver struct {
	db *gorm.DB
}

func (companyDriver *CompanyDriver) BookRegistration(drafts dto.Drafts) error {
	// @todo トランザクション
	// company
	if err := companyDriver.db.Create(&drafts.Company).Error; err != nil {
		return err
	}
	// user
	drafts.User.CompanyID = drafts.Company.ID
	if err := companyDriver.db.Create(&drafts.User).Error; err != nil {
		return err
	}
	// userPassword
	drafts.UserPassword.UserID = drafts.User.ID
	if err := companyDriver.db.Create(&drafts.UserPassword).Error; err != nil {
		return err
	}
	return nil
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
