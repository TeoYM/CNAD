// user_service.go

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

// UserService represents the user service
type UserService struct {
	// Add dependencies here (e.g., database connection)
}

// User represents a user
type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

// NewUserService returns a new user service
func NewUserService() *UserService {
	return &UserService{}
}

// RegisterUser handles user registration
func (s *UserService) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	// Save user to database ( implementation omitted for brevity )
	// ...

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// LoginUser handles user login
func (s *UserService) LoginUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Find user by email ( implementation omitted for brevity )
	// ...

	// Compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte("hashed_password"))
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Generate JWT token ( implementation omitted for brevity )
	// ...

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Login successful")
}

func main() {
	router := mux.NewRouter()
	userService := NewUserService()

	router.HandleFunc("/users", userService.RegisterUser).Methods("POST")
	router.HandleFunc("/login", userService.LoginUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
