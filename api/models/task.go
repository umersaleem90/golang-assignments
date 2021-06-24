package models

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title 				string 	`json:"title" gorm:"not null"`
	Description 	string 	`json:"description" gorm:"not null"`
	UserID				uint		`json:"userID"`		
	User					User		`json:"-"` //omiting it from json as we do not need it in the response
}

func (task *Task) Create(db *gorm.DB) error {
	response := db.Create(&task)
	// No need to preload user table as we don't need it.
	return response.Error
}

func (task *Task) GetAllForUser(db *gorm.DB, ) ([]Task, error) {
	var tasks []Task
	response := db.Where(&Task{UserID: task.UserID}).Find(&tasks)
	return tasks, response.Error
}

func (task *Task) Edit(db *gorm.DB, ) error {
	response := db.Save(&task)
	return response.Error
}

func (task *Task) Delete(db *gorm.DB, ) error {
	response := db.Delete(&task, task.ID)
	return response.Error
}