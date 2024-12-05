package main

import (
	"database/sql"
	"log"
	"net/http"

	"billing-service/handler"
	"billing-service/service"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	// Connect to database
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/assg1")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create billing service
	billingService := service.NewBillingService(db)

	// Create billing handler
	billingHandler := handler.NewBillingHandler(billingService)

	// Create router
	router := mux.NewRouter()

	// Register routes
	router.HandleFunc("/billings", billingHandler.CreateBillingRecord).Methods("POST")
	router.HandleFunc("/billings/{id}", billingHandler.GetBillingRecord).Methods("GET")

	// Start server
	log.Fatal(http.ListenAndServe(":8080", router))
}
