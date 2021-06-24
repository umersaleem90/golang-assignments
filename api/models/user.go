package models

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email 		string `json:"email" gorm:"not null; unique"`
	Password 	string `json:"password" gorm:"not null"`
}

func (user *User) CreateUser(db *gorm.DB) error {
	response := db.Where("email = ?", user.Email).First(&user)
	if response.Error == nil {
		return errors.New("User already existed")
	}
	response = db.Create(&user)
	return response.Error
}

func (user *User) GetUserWithEmailPassword(db *gorm.DB) error {
	response := db.Where(&User{Email: user.Email, Password: user.Password}).First(&user)
	return response.Error
}