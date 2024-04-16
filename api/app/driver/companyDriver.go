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
	err := companyDriver.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&drafts.Company).Error; err != nil {
			return err
		}
		drafts.User.CompanyID = drafts.Company.ID
		if err := tx.Create(&drafts.User).Error; err != nil {
			return err
		}
		drafts.UserPassword.UserID = drafts.User.ID
		if err := tx.Create(&drafts.UserPassword).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
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
