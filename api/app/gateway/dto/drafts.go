package dto

import "go-rest-api/model"

type Drafts struct {
	User         model.User
	Company      model.Company
	UserPassword model.UserPassword
}
