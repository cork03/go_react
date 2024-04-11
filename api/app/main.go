package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"go-rest-api/db"
	"go-rest-api/model"
	"go-rest-api/router"
	"log/slog"
	"os"
)

func main() {
	// todo 環境変数をdevとprdで読み込み方変えるかも。今はdevのみ
	loadErr := godotenv.Load()
	if loadErr != nil {
		println(loadErr)
	}
	//ログの設定 @todo 本番と開発で設定を変えレルようにする
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

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

	e := router.NewRouter()
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", os.Getenv("APP_PORT"))))
}
