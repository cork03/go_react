package gateway

import (
	"go-rest-api/driverBoundary"
	"go-rest-api/gatewayBoundary"
)

type UserGateway struct {
	userDriver driverBoundary.IUserDriver
}

func (u *UserGateway) ExistByEmail(email string) (bool, error) {
	existUser, err := u.userDriver.ExistByEmail(email)
	if err != nil {
		return false, err
	}
	return existUser, nil
}

func NewUserGateway(userDriver driverBoundary.IUserDriver) gatewayBoundary.IUserGateway {
	return &UserGateway{userDriver: userDriver}
}
