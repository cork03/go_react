package driver

import (
	"errors"
	"go-rest-api/driverBoundary"
	"go-rest-api/model"
	"gorm.io/gorm"
)

type UserDriver struct {
	db *gorm.DB
}

func (u *UserDriver) ExistByEmail(email string) (bool, error) {
	user := model.User{}
	err := u.db.Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func NewUserDriver(db *gorm.DB) driverBoundary.IUserDriver {
	return &UserDriver{db: db}
}
