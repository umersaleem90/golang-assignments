package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"todo-backend/api/middleware"
	"todo-backend/api/models"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Router *mux.Router = mux.NewRouter()
var Error error

func createDSN() string {
	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }
	host := os.Getenv("DBHOST")
	user := os.Getenv("DBUSER")
	password := os.Getenv("DBPASSWORD")
  dbname := os.Getenv("DBNAME")
	port := os.Getenv("DBPORT")
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
}

func initializeDatabase() {

	dsn := createDSN()
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
	Router.HandleFunc("/createTaskWithEmail", CreateTaskWithEmail).Methods("POST")
	
	subRouter := Router.PathPrefix("/v1/api").Subrouter()
	subRouter.Use(middleware.ValidateUser)

	subRouter.HandleFunc("/tasks", CreateTask).Methods("POST")
	subRouter.HandleFunc("/tasks", GetAllUserTasks).Methods("GET")
	subRouter.HandleFunc("/tasks/{id:[0-9]+}", EditTask).Methods("PUT")
	subRouter.HandleFunc("/tasks/{id:[0-9]+}", DeleteTask).Methods("DELETE")

	port := fmt.Sprintf(":%s", os.Getenv("SERVERPORT"))
	log.Println("Server starting on port", port)
	log.Fatal(http.ListenAndServe(port, Router))
}

func Start() {
  initializeDatabase()
	initializeRoutes()
}