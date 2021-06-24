package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"todo-backend/api/models"
	"todo-backend/api/response"
	"todo-backend/utilities"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	json.NewDecoder(r.Body).Decode(&user)
	if !utilities.IsValidEmail(user.Email) {
		response.WriteError(w, http.StatusBadRequest, errors.New("invalid email"))
		return
	}

	passwordHash, err := utilities.HashPassword(user.Password)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, errors.New("invalid email"))
		return
	}
	user.Password = passwordHash
	err = user.CreateUser(DB)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, err)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func Login(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	json.NewDecoder(r.Body).Decode(&user)
	if !utilities.IsValidEmail(user.Email) {
		response.WriteError(w, http.StatusBadRequest, errors.New("invalid email"))
		return
	}

	err := user.GetUserWithEmailPassword(DB)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, err)
		return
	}
	token, err := utilities.CreateToken(user.ID)
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, err)
	}
	user.Password = "" //setting password empty for response json
	json.NewEncoder(w).Encode(map[string]interface{}{"user": user, "token": token})
}