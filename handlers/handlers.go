package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var users []User

func HandleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetUsers(w, r)
	case http.MethodPost:
		CreateUser(w, r)
	case http.MethodPut:
		UpdateUser(w, r)
	case http.MethodDelete:
		DeleteUser(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "Invalid request method")
	}
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	// take vector containing all users and put into response
	json.NewEncoder(w).Encode(users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	//request to create a user json
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		// response is the bad request in error case
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Invalid request payload")
		return
	}

	// Perform validation or any other necessary operations
	//adding user creation to users (user is probably create by POST requisition)
	users = append(users, user)

	w.WriteHeader(http.StatusCreated)
	response := map[string]interface{}{
		"message": "User created successfully",
		"user":    user,
	}
	// error or not added to server response
	json.NewEncoder(w).Encode(response)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Implement the logic to update a user
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Implement the logic to delete a user
}
