package common

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"mikro.host/models"
)

var db *gorm.DB

func GetDb(path string) *gorm.DB {
	// return existing connection
	if db != nil {
		return db
	}

	// open new connection
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		panic("could not open db connection")
	}

	// auto migrate models for now
	db.AutoMigrate(&models.User{})

	return db
}
