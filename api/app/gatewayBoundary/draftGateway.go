package gatewayBoundary

import "go-rest-api/domain"

type IDraftGateway interface {
	Create(
		company domain.Company,
		user domain.User,
		userPassword domain.UserPassword,
		mailCertificationId uint,
	) (domain.User, error)
}
