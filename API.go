package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `"json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(people)
	}
}

func main() {
	r := mux.NewRouter()
	people = append(people, Person{ID: "1", Firstname: "Aline", Lastname: "Silva", Address: &Address{City: "City Madrid", State: "State Madrid"}})
	people = append(people, Person{ID: "2", Firstname: "Maria", Lastname: "Silva", Address: &Address{City: "City Getafe", State: "State Madrid"}})
	r.HandleFunc("/contato", GetPeople).Methods("GET")
	r.HandleFunc("/contato/{id}", GetPerson).Methods("GET")
	r.HandleFunc("/contato/{id}", CreatePerson).Methods("POST")
	r.HandleFunc("/contato/{id}", DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(": 8000", nil))
}

// request index page handle
/*
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "This is the index page!")
}
*/

// request contact page handle
/*
func contactHandlerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "This is the contact page!")
}
*/

// request about page handle
/*
func aboutHandlerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "This is the about page!")
}
*/
