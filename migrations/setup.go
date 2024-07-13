package migrations

import (
	"bookApp/db"
	"bookApp/models"
)

func Init() {
	db := db.GetDB()
	db.AutoMigrate(&models.User{}, &models.Book{})
}
