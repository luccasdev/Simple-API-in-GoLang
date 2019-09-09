package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Pets []Pet

type Pet struct {
	Name string `json:"Name"`
	Age int `json:"Age"`
	Description string `json:"Description"`
}

func savePet( w http.ResponseWriter, r *http.Request){
	var pet Pet
	err := json.NewDecoder(r.Body).Decode(&pet)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	fmt.Fprintf(w, "%s is saved with success!", pet.Name)

}

func getPets(w http.ResponseWriter, r *http.Request){
	pets := Pets{
		Pet{Name:"Nino", Age:17, Description:"My Best Dog!"},
	}
	json.NewEncoder(w).Encode(pets)
}

func home(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome home!")
}

func handleRequest() {
	http.HandleFunc("/", home)
	http.HandleFunc("/save/pet", savePet)
	http.HandleFunc("/pets", getPets)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	fmt.Println("API Started")
	handleRequest()
}

