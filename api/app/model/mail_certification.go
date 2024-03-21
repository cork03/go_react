package model

import "time"

type MailCertification struct {
	ID           uint
	Hash         string       `gorm:"not null"`
	Expire       time.Time    `gorm:"not null"`
	DraftUser    DraftUser    `gorm:"constraint:OnDelete:CASCADE"`
	DraftCompany DraftCompany `gorm:"constraint:OnDelete:CASCADE"`
}

type DraftUser struct {
	ID                  uint
	MailCertificationID uint   `gorm:"not null"`
	Email               string `gorm:"not null"`
	Name                string `gorm:"not null"`
	Password            string `gorm:"not null"`
}

type DraftCompany struct {
	ID                  uint
	MailCertificationID uint   `gorm:"not null"`
	Name                string `gorm:"not null"`
	PostalCode          string `gorm:"not null;type:char(7)"`
	Prefecture          string `gorm:"not null;type:varchar(255)"`
	Town                string `gorm:"not null;type:varchar(255)"`
	Area                string `gorm:"not null;type:varchar(255)"`
	Tel                 uint   `gorm:"not null;type:varchar(255)"`
}
