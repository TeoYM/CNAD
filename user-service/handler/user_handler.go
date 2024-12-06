package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"user-service/model"
	"user-service/service"
)

var userService *service.UserService

func init() {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/assg1")
	if err != nil {
		log.Fatal(err)
	}
	userService = service.NewUserService(db)
}

// RegisterUser handles user registration
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdUser, err := userService.RegisterUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var user model.User

	// Check if the request body is JSON
	if r.Header.Get("Content-Type") == "application/json" {
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		// Handle URL-encoded form payload
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user.Email = r.FormValue("email")
		user.Password = r.FormValue("password")
	}

	// Authenticate the user
	token, err := userService.LoginUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
