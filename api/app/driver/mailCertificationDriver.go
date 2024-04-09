package driver

import (
	"errors"
	"go-rest-api/driverBoundary"
	"go-rest-api/model"
	"gorm.io/gorm"
	"log/slog"
	"os"
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

func (mailCertificationDriver *MailCertificationDriver) GetByToken(token string) (*model.MailCertification, error) {
	mailCertification := model.MailCertification{}
	err := mailCertificationDriver.db.Where("token = ?", token).First(&mailCertification).Error
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// @todo slogをdiする
		slog.Info(err.Error())
		return nil, err
	}
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}
	return &mailCertification, nil
}

func NewMailCertificationDriver(db *gorm.DB) driverBoundary.IMailCertificationDriver {
	return &MailCertificationDriver{db}
}
