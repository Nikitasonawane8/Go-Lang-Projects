package main

import (
	"CRUDAPI/usersPckg"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func initializeRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/users", usersPckg.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", usersPckg.GetUser).Methods("GET")
	r.HandleFunc("/users/create", usersPckg.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", usersPckg.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", usersPckg.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func main() {
	usersPckg.InitialMigration()
	initializeRouter()

}
