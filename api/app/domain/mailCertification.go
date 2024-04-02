package domain

import "time"

type MailCertification struct {
	Token  string
	Expire time.Time
}
