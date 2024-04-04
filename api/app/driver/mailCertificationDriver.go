package driver

import (
	"go-rest-api/driverBoundary"
	"go-rest-api/model"
	"gorm.io/gorm"
)

type MailCertificationDriver struct {
	db *gorm.DB
}

func (mailCertificationDriver *MailCertificationDriver) Create(mailCertification model.MailCertification) (model.MailCertification, error) {
	if err := mailCertificationDriver.db.Create(&mailCertification).Error; err != nil {
		return model.MailCertification{}, err
	}
	return mailCertification, nil
}

func NewMailCertificationDriver(db *gorm.DB) driverBoundary.IMailCertificationDriver {
	return &MailCertificationDriver{db}
}
