package gatewayBoundary

import "go-rest-api/domain"

type IMailCertificationGateway interface {
	Create(mailCertification domain.MailCertification) (domain.MailCertification, error)
	GetByToken(token string) (*domain.MailCertification, error)
}
