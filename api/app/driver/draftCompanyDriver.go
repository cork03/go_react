package driver

import (
	"go-rest-api/driverBoundary"
	"go-rest-api/model"
	"gorm.io/gorm"
)

type DraftCompanyDriver struct {
	db *gorm.DB
}

func (draftCompanyDriver *DraftCompanyDriver) Create(company model.DraftCompany) error {
	if err := draftCompanyDriver.db.Create(&company).Error; err != nil {
		return err
	}
	return nil
}

func NewDraftCompanyDriver(db *gorm.DB) driverBoundary.IDraftCompanyDriver {
	return &DraftCompanyDriver{db}
}
