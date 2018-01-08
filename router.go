package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func LoadRoutes() {
	router := mux.NewRouter()
	router.HandleFunc("/version", GetVersion).Methods("GET")

	router.HandleFunc("/cars", GetCars).Methods("GET")
	router.HandleFunc("/cars/{id}", GetCar).Methods("GET")
	//router.HandleFunc("/cars/{id}", CreateCar).Methods("POST")
	router.HandleFunc("/cars/{id}", DeleteCar).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}