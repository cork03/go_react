package model

type User struct {
	ID           uint
	Email        string       `gorm:"not null"`
	Name         string       `gorm:"not null"`
	CompanyID    uint         `gorm:"not null"`
	UserPassword UserPassword `gorm:"constraint:OnDelete:CASCADE"`
}

type UserPassword struct {
	ID       uint
	Password string `gorm:"not null"`
	UserID   uint   `gorm:"not null"`
}
