package models

import (
	"bookApp/db"
	"bookApp/validations"
	"fmt"
)

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Email    string `json:"email" gorm:"unique"`
	Fullname string `json:"fullname"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role" gorm:"default:'user'"`	
}

func (u User) GetUser(input *validations.LoginInput) (*User, error) {
	db := db.GetDB()
	var user User
	if err := db.Where("email = ?", input.Email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func (u User) CreateUser(input *validations.CreateUserInput) (*User, error) {
	db := db.GetDB()
	user := User{Fullname: input.Fullname, Email: input.Email, Username: input.Username, Password: input.Password}
	if err := db.Create(&user).Error; err != nil {
		if err.Error() == "UNIQUE constraint failed: users.email" {
			return nil, fmt.Errorf("email already exists")
		}
	}
	return &user, nil
}
