package driverBoundary

import "go-rest-api/model"

type IDraftDriver interface {
	Create(
		company model.DraftCompany,
		user model.DraftUser,
		userPassword model.DraftUserPassword,
	) (model.DraftUser, error)
}
