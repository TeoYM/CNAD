package service

import (
	"database/sql"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	"user-service/model"
)

// UserService represents the user service
type UserService struct {
	db *sql.DB
}

// NewUserService returns a new user service
func NewUserService(db *sql.DB) *UserService {
	return &UserService{db: db}
}

// RegisterUser registers a new user
func (s *UserService) RegisterUser(user model.User) (model.User, error) {
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return user, err
	}
	user.Password = string(hashedPassword)

	// Save user to database
	_, err = s.db.Exec("INSERT INTO Users (Email, PasswordHash, Name, CreatedAt) VALUES (?, ?, ?, ?)", user.Email, user.Password, user.Name, time.Now())
	if err != nil {
		return user, err
	}

	return user, nil
}

// LoginUser logs in a user and returns a JWT token
func (s *UserService) LoginUser(user model.User) (string, error) {
	// Find user by email
	var hashedPassword string
	err := s.db.QueryRow("SELECT PasswordHash FROM Users WHERE Email = ?", user.Email).Scan(&hashedPassword)
	if err != nil {
		return "", err
	}

	// Compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.Password))
	if err != nil {
		return "", err
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
	})
	tokenString, err := token.SignedString([]byte("secretkey"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
