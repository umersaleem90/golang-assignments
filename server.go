package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var router *mux.Router = mux.NewRouter()

func setupServer() {
	
	router.HandleFunc("/v1/api/products", GetProducts).Methods("GET")
	router.HandleFunc("/v1/api/products/{id}", GetProduct).Methods("GET")
	router.HandleFunc("/v1/api/products", CreateProduct).Methods("POST")
	router.HandleFunc("/v1/api/products/{id}", UpdateProduct).Methods("PUT")
	router.HandleFunc("/v1/api/products/{id}", DeleteProduct).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", router))
	
}