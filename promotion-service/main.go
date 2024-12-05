package main

import (
	"database/sql"
	"log"
	"net/http"

	"promotion-service/handler"
	"promotion-service/service"

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

	// Create promotion service
	promotionService := service.NewPromotionService(db)

	// Create promotion handler
	promotionHandler := handler.NewPromotionHandler(promotionService)

	// Create router
	router := mux.NewRouter()

	// Register routes
	router.HandleFunc("/promotions", promotionHandler.CreatePromotion).Methods("POST")
	router.HandleFunc("/promotions/{id}", promotionHandler.GetPromotion).Methods("GET")
	router.HandleFunc("/promotions/{id}", promotionHandler.UpdatePromotion).Methods("PUT")
	router.HandleFunc("/promotions/{id}", promotionHandler.DeletePromotion).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":8080", router))
}
