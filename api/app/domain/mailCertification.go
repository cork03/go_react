package domain

import "time"

type MailCertification struct {
	ID     uint
	Token  string
	Expire time.Time
}

type Drafts struct {
	MailCertification MailCertification
	DraftUser         User
	DraftCompany      Company
	DraftUserPassword DraftUserPassword
}
