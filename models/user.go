package models

import (
	"bookApp/db"
)

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Fullname string `json:"fullname"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u User) GetUser(input *LoginInput) (*User, error) {
	db := db.GetDB()
	var user User
	if err := db.Where("username = ?", input.Username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func (u User) CreateUser(input *CreatedUser) (*User, error) {
	db := db.GetDB()
	user := User{Username: input.Username, Password: input.Password}
	if err := db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type RegisterInput struct {
	Fullname       string `json:"fullname" binding:"required"`
	Username       string `json:"username" binding:"required"`
	Password       string `json:"password" binding:"required"`
	ConfirmPasword string `json:"confirm_password" binding:"required"`
}
type CreatedUser struct {
	Fullname string `json:"fullname" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
