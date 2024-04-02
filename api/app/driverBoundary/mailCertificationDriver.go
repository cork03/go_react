package driverBoundary

import "go-rest-api/model"

type IMailCertificationDriver interface {
	Create(mailCertification model.MailCertification) error
}
