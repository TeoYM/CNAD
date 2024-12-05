package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	service "membership-service/handler"
	membershipSvc "membership-service/service"
)

func main() {
	// Connect to database
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/assg1")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create membership service
	membershipService := membershipSvc.NewMembershipService(db)

	// Create membership handler
	membershipHandler := service.NewMembershipHandler(membershipService)

	// Create router
	router := mux.NewRouter()

	// Register routes
	router.HandleFunc("/memberships", membershipHandler.CreateMembership).Methods("POST")
	router.HandleFunc("/memberships/{user_id}", membershipHandler.GetMembership).Methods("GET")

	// Start server
	log.Fatal(http.ListenAndServe(":8080", router))
}
