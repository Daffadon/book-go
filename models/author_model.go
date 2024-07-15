package models

import (
	"bookApp/db"
	"bookApp/validations"
	"fmt"
)

type Author struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Email    string `json:"email" gorm:"unique"`
	Fullname string `json:"fullname"`
	Books    []Book `json:"books" gorm:"foreignKey:AuthorID"`
}

func (a Author) GetAuthors() ([]Author, error) {
	db := db.GetDB()
	var arrAuthors []Author
	if err := db.Find(&arrAuthors).Error; err != nil {
		if err.Error() == "record not found" {
			return nil, fmt.Errorf("record not found")
		}
	}
	return arrAuthors, nil
}

func (a Author) CreateAuthors(input *validations.CreateAuthorInput) (*Author, error) {
	db := db.GetDB()
	author := Author{Fullname: input.Fullname, Email: input.Email}
	if err := db.Create(&author).Error; err != nil {
		if err.Error() == "UNIQUE constraint failed: authors.email" {
			return nil, fmt.Errorf("email already exists")
		}
	}
	return &author, nil
}
