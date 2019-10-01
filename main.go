package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var petsArray []Pet

type Pet struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"Name"`
	Age         int    `json:"Age"`
	Description string `json:"Description"`
}

func savePet(w http.ResponseWriter, r *http.Request) {
	var pet Pet
	_ = json.NewDecoder(r.Body).Decode(&pet)
	petsArray = append(petsArray, pet)
	json.NewEncoder(w).Encode(petsArray)
	fmt.Fprintf(w, "%s is saved with success!", pet.Name)

}

func getAllPets(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(petsArray)
}

func getPet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range petsArray {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Pet{})

}

func deletePet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range petsArray {
		if item.ID == params["id"] {
			petsArray = append(petsArray[:index], petsArray[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(petsArray)
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

	petsArray = append(petsArray, Pet{ID: "1", Name: "John", Age: 12, Description: "Border Collie"})
	petsArray = append(petsArray, Pet{ID: "2", Name: "Toby", Age: 12, Description: "Labrador"})

	handleRequest()
}
