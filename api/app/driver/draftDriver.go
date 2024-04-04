package driver

import (
	"go-rest-api/driverBoundary"
	"go-rest-api/model"
	"gorm.io/gorm"
)

type DraftDriver struct {
	db *gorm.DB
}

func (draftDriver *DraftDriver) Create(
	company model.DraftCompany,
	user model.DraftUser,
	userPassword model.DraftUserPassword,
) (model.DraftUser, error) {
	if err := draftDriver.db.Create(&company).Error; err != nil {
		return model.DraftUser{}, err
	}
	if err := draftDriver.db.Create(&user).Error; err != nil {
		return model.DraftUser{}, err
	}
	if err := draftDriver.db.Create(&userPassword).Error; err != nil {
		return model.DraftUser{}, err
	}
	return user, nil
}

func NewDraftDriver(db *gorm.DB) driverBoundary.IDraftDriver {
	return &DraftDriver{db}
}
