package config

import (
	"fmt"
	"my-gram/model"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func DatabaseConnect() {
	dsn := fmt.Sprintf(
		// "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("TZ"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.Debug().AutoMigrate(
		model.User{},
		model.SocialMedia{},
		model.Photo{},
		model.Comment{},
	)

	Db = db

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connection success")
	}
}
