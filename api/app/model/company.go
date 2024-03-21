package model

type Company struct {
	ID         uint
	Name       string `gorm:"not null"`
	PostalCode string `gorm:"not null;type:char(7)"`
	Prefecture string `gorm:"not null;type:varchar(255)"`
	Town       string `gorm:"not null;type:varchar(255)"`
	Area       string `gorm:"not null;type:varchar(255)"`
	Tel        uint   `gorm:"not null;type:varchar(255)"`
	User       []User
}
