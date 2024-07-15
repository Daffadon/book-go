package models

import (
	"bookApp/db"
	"bookApp/validations"
	"fmt"
)

type Book struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	AuthorID    uint   `json:"author_id" gorm:"not null"`
	Title       string `json:"title" gorm:"not null"`
	Description string `json:"description"`
	Genre       string `json:"genre" gorm:"not null"`
	NumPages    int    `json:"num_pages" gorm:"not null"`
	Languages   string `json:"languages" gorm:"not null"`
	Stock       int    `json:"stock" gorm:"not null"`
	Price       int    `json:"price" gorm:"not null"`
	Author      Author `json:"-" gorm:"foreignKey:AuthorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (b Book) GetBooks() ([]Book, error) {
	db := db.GetDB()
	var arrBook []Book
	if err := db.Find(&arrBook).Error; err != nil {
		if err.Error() == "record not found" {
			return nil, fmt.Errorf("record not found")
		}
	}
	return arrBook, nil
}

func (b Book) GetBookByID(id string) (*Book, error) {
	db := db.GetDB()
	var bookRetrieved Book
	if err := db.Where("id = ?", id).First(&bookRetrieved).Error; err != nil {
		return nil, err
	}
	return &bookRetrieved, nil
}

func (b Book) CreateBook(input *validations.CreateBookServiceInput) (*Book, error) {
	db := db.GetDB()
	book := Book{Title: input.Title, Description: input.Description, Genre: input.Genre, NumPages: input.NumPages, Languages: input.Languages, Stock: input.Stock, Price: input.Price, AuthorID: input.AuthorID}
	if err := db.Create(&book).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func (b Book) UpdateBook(id string, input *validations.UpdateBookServiceInput) (*Book, error) {
	db := db.GetDB()
	var updatedBook Book
	if err := db.Where("id = ?", id).First(&updatedBook).Error; err != nil {
		if err.Error() == "record not found" {
			return nil, fmt.Errorf("record not found")
		}
	}
	if err := db.Model(&updatedBook).Updates(&input).Error; err != nil {
		return nil, err
	}
	return &updatedBook, nil
}

func (b Book) DeleteBook(id string) (bool, error) {
	db := db.GetDB()
	var deletedBook Book
	if err := db.Where("id = ", id).First(&deletedBook).Error; err != nil {
		return false, err
	}
	if err := db.Delete(&deletedBook).Error; err != nil {
		return false, err
	}
	return true, nil
}
