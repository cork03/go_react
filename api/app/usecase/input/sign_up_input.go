package input

import "go-rest-api/domain"

type SignUpInput struct {
	Company      domain.Company
	User         domain.User
	UserPassword domain.UserPassword
}
