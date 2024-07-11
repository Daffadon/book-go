package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {

	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}
	db = database
}

func GetDB() *gorm.DB {
	return db
}
