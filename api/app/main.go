package main

import (
	"fmt"
	"go-rest-api/db"
	"go-rest-api/model"
)

func main() {
	db := db.Main()
	err := db.AutoMigrate(
		&model.Company{},
		&model.User{},
		&model.UserPassword{},
		&model.MailCertification{},
		&model.DraftCompany{},
		&model.DraftUser{},
	)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Hello, World!")
}
