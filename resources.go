package main

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

// Return version
func GetVersion(writer http.ResponseWriter, _ *http.Request) {
	json.NewEncoder(writer).Encode(Version{"0.1"})
}

// Get all cars
func GetCars(writer http.ResponseWriter, _ *http.Request) {
	json.NewEncoder(writer).Encode(cars)
}

// Get single car
func GetCar(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	for _, item := range cars {
		if item.ID == params["id"] {
			json.NewEncoder(writer).Encode(item)
			return
		}
	}
	json.NewEncoder(writer).Encode(&Car{})
}

// Create a car
func CreateCar(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	var car Car
	_ = json.NewDecoder(request.Body).Decode(&car)
	car.ID = params["id"]
	cars = append(cars, car)
	json.NewEncoder(writer).Encode(cars)
}

// Delete a car
func DeleteCar(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	for index, item := range cars {
		if item.ID == params["id"] {
			cars = append(cars[:index], cars[index+1:]...)
			break
		}
		json.NewEncoder(writer).Encode(cars)
	}
}
