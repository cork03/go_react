package driver

import (
	"go-rest-api/driverBoundary"
	"go-rest-api/model"
	"gorm.io/gorm"
)

type MailCertificationDriver struct {
	db *gorm.DB
}

func (mailCertificationDriver *MailCertificationDriver) Create(mailCertification model.MailCertification) error {
	if err := mailCertificationDriver.db.Create(&mailCertification).Error; err != nil {
		return err
	}
	return nil
}

func NewMailCertificationDriver(db *gorm.DB) driverBoundary.IMailCertificationDriver {
	return &MailCertificationDriver{db}
}
