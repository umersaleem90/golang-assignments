package controllers

import (
	"fmt"
	"log"
	"net/http"
	"todo-backend/api/middleware"
	"todo-backend/api/models"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Router *mux.Router = mux.NewRouter()
var Error error

const dsn = "host=localhost user=umersaleem password=1122 dbname=todoDB port=8000 sslmode=disable"

func initializeDatabase() {
	DB, Error = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if Error != nil {
		fmt.Println(Error.Error())
		panic("Cannot connect to DB")
	}
	DB.AutoMigrate(&models.User{}, &models.Task{})
}

func initializeRoutes() {
	
	Router.Use(middleware.SetContentTypeMiddleware)
	Router.HandleFunc("/signUp", SignUp).Methods("POST")
	Router.HandleFunc("/login", Login).Methods("POST")
	
	subRouter := Router.PathPrefix("/v1/api").Subrouter()
	subRouter.Use(middleware.ValidateUser)

	subRouter.HandleFunc("/tasks", CreateTask).Methods("POST")
	subRouter.HandleFunc("/tasks", GetAllUserTasks).Methods("GET")
	subRouter.HandleFunc("/tasks/{id:[0-9]+}", EditTask).Methods("PUT")
	subRouter.HandleFunc("/tasks/{id:[0-9]+}", DeleteTask).Methods("DELETE")

	log.Println("Server starting on port 3000")
	log.Fatal(http.ListenAndServe(":3000", Router))
	
}

func Start() {
  initializeDatabase()
	initializeRoutes()
}