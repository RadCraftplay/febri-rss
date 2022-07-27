package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

const (
	host     = "localhost"
	port     = 5432
	username = "febri"
	password = "e3b3440172abd558bbd535eefbc36512a0a574b5360bc4d59d14fc9404dc8bbc"
	dbname   = "febri"
)

func ConnectDatabase() {
	database, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, username, password, dbname, port),
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Feed{})
	database.AutoMigrate(&SentItems{})

	DB = database
}
