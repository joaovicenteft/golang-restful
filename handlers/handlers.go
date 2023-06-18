package handlers

import (
	"encoding/json"
	//"fmt"
	//"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}


var Users []User

func GetUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for _, user := range Users{
		if user.ID == id {
			json.NewEncoder(w).Encode(user)
			return
		}
	}

	json.NewEncoder(w).Encode("User not found")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	json.NewDecoder(r.Body).Decode(&newUser)
	Users= append(Users, newUser)
	json.NewEncoder(w).Encode(Users)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for index, user := range Users{
		if user.ID == id {
			Users= append(Users[:index], Users[index+1:]...) // Remove the old user
			var updatedUser User
			json.NewDecoder(r.Body).Decode(&updatedUser)
			updatedUser.ID = id
			Users= append(Users, updatedUser) // Add the updated user
			json.NewEncoder(w).Encode(updatedUser)
			return
		}
	}

	json.NewEncoder(w).Encode("User not found")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for index, user := range Users{
		if user.ID == id {
			Users= append(Users[:index], Users[index+1:]...) // Remove the user from the slice
			break
		}
	}

	json.NewEncoder(w).Encode(Users)
}