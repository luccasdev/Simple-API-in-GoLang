package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var pets []Pet

type Pet struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"Name"`
	Age         int    `json:"Age"`
	Description string `json:"Description"`
}

func savePet(w http.ResponseWriter, r *http.Request) {
	var pet Pet
	err := json.NewDecoder(r.Body).Decode(&pet)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	fmt.Fprintf(w, "%s is saved with success!", pet.Name)

}

func getAllPets(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(pets)
}

func getPet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func deletePet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range pets {
		if item.ID == params["id"] {
			pets = append(pets[:index], pets[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(pets)
	}
}

func handleRequest() {
	//Routes
	router := mux.NewRouter()

	router.HandleFunc("/pet/{id}", getPet).Methods("GET")
	router.HandleFunc("/pet", getAllPets).Methods("GET")
	router.HandleFunc("/pet", savePet).Methods("POST")
	router.HandleFunc("/pet/{id}", deletePet).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	fmt.Println("API Started")
	handleRequest()

	pets = append(pets, Pet{ID: "1", Name: "John", Age: 12, Description: "Border Collie"})
	pets = append(pets, Pet{ID: "2", Name: "Toby", Age: 12, Description: "Labrador"})
}
