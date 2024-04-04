package domain

import "time"

type MailCertification struct {
	ID     uint
	Token  string
	Expire time.Time
}
