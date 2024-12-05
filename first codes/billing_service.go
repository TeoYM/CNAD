package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// BillingService represents the billing service
type BillingService struct {
	db *sql.DB
}

// NewBillingService returns a new billing service
func NewBillingService(db *sql.DB) *BillingService {
	return &BillingService{db: db}
}

// Billing represents a billing record
type Billing struct {
	ID           string    `json:"id"`
	UserID       string    `json:"user_id"`
	VehicleID    string    `json:"vehicle_id"`
	RentalPeriod string    `json:"rental_period"`
	TotalCost    float64   `json:"total_cost"`
	CreatedAt    time.Time `json:"created_at"`
}

// CreateBillingRecord handles billing record creation
func (s *BillingService) CreateBillingRecord(w http.ResponseWriter, r *http.Request) {
	var billing Billing
	err := json.NewDecoder(r.Body).Decode(&billing)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Calculate total cost based on rental period and vehicle type
	// ( implementation omitted for brevity )
	// ...

	// Save billing record to database
	_, err = s.db.Exec("INSERT INTO Billings (UserID, VehicleID, RentalPeriod, TotalCost, CreatedAt) VALUES (?, ?, ?, ?, ?)", billing.UserID, billing.VehicleID, billing.RentalPeriod, billing.TotalCost, time.Now())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(billing)
}

// GetBillingRecord handles billing record retrieval
func (s *BillingService) GetBillingRecord(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	billingID := params["id"]

	// Find billing record by ID
	var billing Billing
	err := s.db.QueryRow("SELECT * FROM Billings WHERE ID = ?", billingID).Scan(&billing.ID, &billing.UserID, &billing.VehicleID, &billing.RentalPeriod, &billing.TotalCost, &billing.CreatedAt)
	if err != nil {
		http.Error(w, "Billing record not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(billing)
}

// UpdateBillingRecord handles billing record updates
func (s *BillingService) UpdateBillingRecord(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	billingID := params["id"]

	var billing Billing
	err := json.NewDecoder(r.Body).Decode(&billing)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update billing record in database
	_, err = s.db.Exec("UPDATE Billings SET UserID = ?, VehicleID = ?, RentalPeriod = ?, TotalCost = ? WHERE ID = ?", billing.UserID, billing.VehicleID, billing.RentalPeriod, billing.TotalCost, billingID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(billing)
}

// DeleteBillingRecord handles billing record deletion
func (s *BillingService) DeleteBillingRecord(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	billingID := params["id"]

	// Delete billing record from database
	_, err := s.db.Exec("DELETE FROM Billings WHERE ID = ?", billingID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func main() {
	// Connect to database
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/my_db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create billing service
	billingService := NewBillingService(db)

	// Create router
	router := mux.NewRouter()

	// Register routes
	router.HandleFunc("/billings", billingService.CreateBillingRecord).Methods("POST")
	router.HandleFunc("/billings/{id}", billingService.GetBillingRecord).Methods("GET")
	router.HandleFunc("/billings/{id}", billingService.UpdateBillingRecord).Methods("PUT")
	router.HandleFunc("/billings/{id}", billingService.DeleteBillingRecord).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":8082", router))
}
