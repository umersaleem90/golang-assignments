package models

import (
	"errors"
	"todo-backend/utilities"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email 				string 	`json:"email" gorm:"not null; unique"`
	Password 			string 	`json:"password" gorm:"not null"`
	IsDummyUser		bool 		`json:"isDummyUser" gorm:"default:false"`
}

func (user *User) CreateUser(db *gorm.DB) error {
	var password string = user.Password
	response := db.Where(&User{Email: user.Email}).First(&user)
	if response.Error == nil {
		if user.IsDummyUser {
			user.IsDummyUser = false
			user.Password = password
			db.Save(&user)
			return nil
		}
		return errors.New("User already existed")
	}
	response = db.Create(&user)
	return response.Error
}

func (user *User) GetUserWithEmailPassword(db *gorm.DB) error {
	var password string = user.Password
	response := db.Where(&User{Email: user.Email}).First(&user)
	if response.Error != nil {
		return response.Error
	}
	err := utilities.ComparePasswordHash(password, user.Password)
	return err
}