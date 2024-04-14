package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"go-rest-api/db"
	"go-rest-api/model"
)

func main() {
	loadErr := godotenv.Load()
	if loadErr != nil {
		println(loadErr)
	}

	db := db.Main()
	err := db.AutoMigrate(
		&model.Company{},
		&model.User{},
		&model.UserPassword{},
		&model.MailCertification{},
		&model.DraftCompany{},
		&model.DraftUser{},
		&model.DraftUserPassword{},
	)
	if err != nil {
		fmt.Println(err)
	}
}
