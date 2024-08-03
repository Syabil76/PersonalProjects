package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Car struct {
	ID      string `json: "ID"`
	Make    string `json: Make`
	Model   string `json: Model`
	Year    int    `json: year`
	Edition string `json: Edition`
	Class   string `json: Class`
	Brand   *Brand `json: Brand`
}

type Brand struct {
	Name        string `json: Name`
	Nationality string `json: Nationality`
}

var cars []Car

func READALL(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cars)
}

func DELETE(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range cars {
		if item.ID == params["id"] {
			cars = append(cars[:index], cars[index+1:]...)
			break
		}
	}
	json.NewEncoder().Encode(cars)
}

func READ(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range cars {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		} else {
			fmt.Println("Car ID not found!")
		}
	}
}

func CREATE(w http.ResponseWriter, r *http.Response) {
	w.Header().Set("Content-Type", "applications/json")
	var car Car
	_ = json.NewDecoder(r.Body).Decode(&car)
	car.ID = strconv.Itoa(rand.Intn(1000))
	cars = append(cars, car)
	json.NewEncoder(w).Encode(car)
}

func UPDATE(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range cars {
		if item.ID == params["id"] {
			cars = append(cars[:index], cars[index+1:]...)
			var car Car
			cars = append(cars, car)
			_ = json.NewDecoder(r.Body).Decode(&car)
			car.ID = params["id"]
			cars = append(cars, car)
			json.NewEncoder(w).Encode(cars)
		}
	}
}

func main() {
	r := mux.NewRouter()

	// writing the handles for the ROUTES for logic

	cars = append(cars, Car{ID: "1000", Make: "Mazda", Model: "RX-7", Year: 1990, Edition: "FC3S", Class: "Supercar", Brand: &Brand{Name: "Mazda", Nationality: "Japanese"}})
	r.HandleFunc("/cars", READALL).Methods("GET")
	r.HandleFunc("/cars/{id}", READ).Methods("GET")
	r.HandleFunc("/cars", CREATE).Methods("POST")
	r.HandleFunc("/cars/id", UPDATE).Methods("PUT")
	r.HandleFunc("/cars/{id}", DELETE).Methods("DELETE")

	fmt.Print("Starting server on Port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))

}
