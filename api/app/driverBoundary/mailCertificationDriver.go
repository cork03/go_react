package driverBoundary

import "go-rest-api/model"

type IMailCertificationDriver interface {
	Create(mailCertification model.MailCertification) (model.MailCertification, error)
	GetByTokenWithDrafts(token string) (*model.MailCertification, error)
}
