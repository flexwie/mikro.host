package common

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"mikro.host/models"
	"os"
)

var db *gorm.DB

func GetDb(path *string) *gorm.DB {
	// return existing connection
	if db != nil {
		return db
	}

	// open new connection
	var err error
	if path == nil {
		dbHost := os.Getenv("DB_HOST")
		dbUser := os.Getenv("DB_USER")
		dbPwd := os.Getenv("DB_PWD")
		dbName := os.Getenv("DB_NAME")

		db, err = gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disabled", dbHost, dbUser, dbPwd, dbName)), &gorm.Config{})
	} else {
		db, err = gorm.Open(sqlite.Open(*path), &gorm.Config{})
	}

	if err != nil {
		panic("could not open db connection")
	}

	// auto migrate models for now
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		panic("could not migrate db")
	}

	return db
}
