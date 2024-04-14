package main

import (
	"fmt"
	"github.com/joho/godotenv"
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
	//ログの設定 @todo 本番と開発で設定を変えられるようにする
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	e := router.NewRouter()
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", os.Getenv("APP_PORT"))))
}
