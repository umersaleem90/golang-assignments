package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"todo-backend/api/models"
	"todo-backend/api/response"
	"todo-backend/utilities"

	"github.com/gorilla/mux"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	task := models.Task{}
	task.UserID = r.Context().Value("userID").(uint)
	json.NewDecoder(r.Body).Decode(&task)
	err := task.Create(DB)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, err)
		return
	}
	json.NewEncoder(w).Encode(task)
}

func GetAllUserTasks(w http.ResponseWriter, r *http.Request) { 
	task := models.Task{}
	task.UserID = r.Context().Value("userID").(uint)
	tasks, err := task.GetAllForUser(DB)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, err)
		return
	}
	json.NewEncoder(w).Encode(tasks)
}

func EditTask(w http.ResponseWriter, r *http.Request) { 
	task := models.Task{}
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 32)
	task.ID = uint(id)
	DB.First(&task, task.ID)
	json.NewDecoder(r.Body).Decode(&task)
	err := task.Edit(DB)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, err)
		return
	}
	json.NewEncoder(w).Encode(task)
}


func DeleteTask(w http.ResponseWriter, r *http.Request) { 
	task := models.Task{}
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 32)
	task.ID = uint(id)
	json.NewDecoder(r.Body).Decode(&task)
	err := task.Delete(DB)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, err)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Product deleted"})
}

func CreateTaskWithEmail(w http.ResponseWriter, r *http.Request) { 
	task := models.Task{}
	email := r.URL.Query().Get("email")

	if !utilities.IsValidEmail(email) {
		response.WriteError(w, http.StatusBadRequest, errors.New("invalid email"))
		return
	}
	var user models.User
	user.Email = email
	resp := DB.Where(&models.User{Email: email}).First(&user)

	if resp.Error != nil {
		user.IsDummyUser = true
		err := user.CreateUser(DB)
		if err != nil {
			response.WriteError(w, http.StatusInternalServerError, errors.New("invalid email"))
			return
		}
	}

	task.UserID = user.ID
	json.NewDecoder(r.Body).Decode(&task)
	err := task.Create(DB)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, err)
		return
	}
	json.NewEncoder(w).Encode(task)
}

