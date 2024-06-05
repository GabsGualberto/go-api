package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var users []User

func main() {
	//TODO: Buscar dados do db
	users = append(users, User{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
	users = append(users, User{ID: "2", Firstname: "Matias", Lastname: "Rodigres", Address: &Address{City: "São Pulo", State: "SP"}})
	users = append(users, User{ID: "3", Firstname: "Walati", Lastname: "Cazuza", Address: &Address{City: "Cairas", State: "SP"}})

	router := mux.NewRouter()
	router.HandleFunc("/users", GetUsers).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}
