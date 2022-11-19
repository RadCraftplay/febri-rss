package models

import (
	"fmt"

	"febri-rss/common"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(configuration common.FebriRssConfiguration) {
	database, err := gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
			configuration.Postgres.Host,
			configuration.Postgres.Username,
			configuration.Postgres.Password,
			configuration.Postgres.DbName,
			configuration.Postgres.Port),
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	db, err := database.DB()
	if err != nil {
		panic("Error getting databse out of *gorm.DB!")
	}

	err = db.Ping()
	if err != nil {
		panic("Failed to ping the database!")
	}

	database.AutoMigrate(&Feed{})
	database.AutoMigrate(&SentItems{})

	DB = database
}
