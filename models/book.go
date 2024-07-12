package models

import (
	"bookApp/db"
)

type Book struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func (b Book) GetBookByID(id string) (*Book, error) {
	db := db.GetDB()
	var bookRetrieved Book
	if err := db.Where("id = ?", id).First(&bookRetrieved).Error; err != nil {
		return nil, err
	}
	return &bookRetrieved, nil
}

func (b Book) CreateBook(input *CreateBookInput) (*Book, error) {
	db := db.GetDB()
	book := Book{Title: input.Title, Author: input.Author}
	if err := db.Create(&book).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func (b Book) UpdateBook(id string, input *UpdateBookInput) (*Book, error) {
	db := db.GetDB()
	var updatedBook Book
	if err := db.Where("id = ?", id).First(&updatedBook).Error; err != nil {
		return nil, err
	}
	if err := db.Model(&updatedBook).Updates(&input).Error; err != nil {
		return nil, err
	}
	return &updatedBook, nil
}

func (b Book) DeleteBook(id string) (bool, error) {
	db := db.GetDB()
	var deletedBook Book
	if err := db.Where("id = ?", id).First(&deletedBook).Error; err != nil {
		return false, err
	}
	if err := db.Delete(&deletedBook).Error; err != nil {
		return false, err
	}
	return true, nil
}

type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}
