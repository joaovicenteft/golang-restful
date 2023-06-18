package main

import (
	"log"
	"net/http"
	"restfulImp/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	
	// Add sample data
	handlers.Users = append(handlers.Users, handlers.User{ID: 1, Username: "john", Email: "john@example.com"})
	handlers.Users = append(handlers.Users, handlers.User{ID: 2, Username: "jane", Email: "jane@example.com"})


	// Routes
	router.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
	router.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

