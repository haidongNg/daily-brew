package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", AppConfig.DBHost, AppConfig.DBUser, AppConfig.DBPassword, AppConfig.DBName)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to DB: ", err)
	}

	log.Println("Connected to DB")
}
