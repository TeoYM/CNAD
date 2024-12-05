package main

import (
	"log"
	"net/http"

	"vehicle-service/handler"

	"github.com/gorilla/mux"
)

func main() {
	// Create router
	router := mux.NewRouter()

	// Register routes
	router.HandleFunc("/vehicles", handler.CreateVehicle).Methods("POST")
	router.HandleFunc("/vehicles/{id}", handler.GetVehicle).Methods("GET")
	router.HandleFunc("/vehicles/{id}", handler.UpdateVehicle).Methods("PUT")
	router.HandleFunc("/vehicles/{id}", handler.DeleteVehicle).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":8080", router))
}
