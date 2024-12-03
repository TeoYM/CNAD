package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

// UserService represents the user service
type UserService struct {
	db *sql.DB
}

// NewUserService returns a new user service
func NewUserService(db *sql.DB) *UserService {
	return &UserService{db: db}
}

// User represents a user
type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
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

	// Save user to database
	_, err = s.db.Exec("INSERT INTO Users (Email, PasswordHash, Name, CreatedAt) VALUES (?, ?, ?, ?)", user.Email, user.Password, user.Name, time.Now())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

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

	// Find user by email
	var hashedPassword string
	err = s.db.QueryRow("SELECT PasswordHash FROM Users WHERE Email = ?", user.Email).Scan(&hashedPassword)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.Password))
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
	// Connect to database
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/my_db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create user service
	userService := NewUserService(db)

	// Create router
	router := mux.NewRouter()

	// Register routes
	router.HandleFunc("/users", userService.RegisterUser).Methods("POST")
	router.HandleFunc("/login", userService.LoginUser).Methods("POST")

	// Start server
	log.Fatal(http.ListenAndServe(":8080", router))
}
