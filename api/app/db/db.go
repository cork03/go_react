package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func Main() *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PW"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DB"),
	)
	// @ todo devではinfoに吐く本番だと特定のログに吐く
	gormLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: 2 * time.Second,
			LogLevel:      logger.Info,
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                 gormLogger,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		println(err)
	}
	return db
}
