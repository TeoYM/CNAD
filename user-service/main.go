package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"database/sql"
	"user-service/handler"

	"github.com/gorilla/mux"
)

func main() {
	// Connect to database
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/assg1")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create router
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "login.html")
	})

	// Register routes
	router.HandleFunc("/users", handler.RegisterUser).Methods("POST")
	router.HandleFunc("/login", handler.LoginUser).Methods("POST")

	// Start server
	log.Fatal(http.ListenAndServe(":8080", router))
}
