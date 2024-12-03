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

// VehicleService represents the vehicle service
type VehicleService struct {
	db *sql.DB
}

// NewVehicleService returns a new vehicle service
func NewVehicleService(db *sql.DB) *VehicleService {
	return &VehicleService{db: db}
}

// Vehicle represents a vehicle
type Vehicle struct {
	ID           string    `json:"id"`
	LicensePlate string    `json:"license_plate"`
	VehicleType  string    `json:"vehicle_type"`
	Availability bool      `json:"availability"`
	CreatedAt    time.Time `json:"created_at"`
}

// CreateVehicle handles vehicle creation
func (s *VehicleService) CreateVehicle(w http.ResponseWriter, r *http.Request) {
	var vehicle Vehicle
	err := json.NewDecoder(r.Body).Decode(&vehicle)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Save vehicle to database
	_, err = s.db.Exec("INSERT INTO Vehicles (LicensePlate, VehicleType, Availability, CreatedAt) VALUES (?, ?, ?, ?)", vehicle.LicensePlate, vehicle.VehicleType, vehicle.Availability, time.Now())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(vehicle)
}

// GetVehicle handles vehicle retrieval
func (s *VehicleService) GetVehicle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	vehicleID := params["id"]

	// Find vehicle by ID
	var vehicle Vehicle
	err := s.db.QueryRow("SELECT * FROM Vehicles WHERE ID = ?", vehicleID).Scan(&vehicle.ID, &vehicle.LicensePlate, &vehicle.VehicleType, &vehicle.Availability, &vehicle.CreatedAt)
	if err != nil {
		http.Error(w, "Vehicle not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vehicle)
}

// UpdateVehicle handles vehicle updates
func (s *VehicleService) UpdateVehicle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	vehicleID := params["id"]

	var vehicle Vehicle
	err := json.NewDecoder(r.Body).Decode(&vehicle)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update vehicle in database
	_, err = s.db.Exec("UPDATE Vehicles SET LicensePlate = ?, VehicleType = ?, Availability = ? WHERE ID = ?", vehicle.LicensePlate, vehicle.VehicleType, vehicle.Availability, vehicleID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vehicle)
}

// DeleteVehicle handles vehicle deletion
func (s *VehicleService) DeleteVehicle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	vehicleID := params["id"]

	// Delete vehicle from database
	_, err := s.db.Exec("DELETE FROM Vehicles WHERE ID = ?", vehicleID)
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

	// Create vehicle service
	vehicleService := NewVehicleService(db)

	// Create router
	router := mux.NewRouter()

	// Register routes
	router.HandleFunc("/vehicles", vehicleService.CreateVehicle).Methods("POST")
	router.HandleFunc("/vehicles/{id}", vehicleService.GetVehicle).Methods("GET")
	router.HandleFunc("/vehicles/{id}", vehicleService.UpdateVehicle).Methods("PUT")
	router.HandleFunc("/vehicles/{id}", vehicleService.DeleteVehicle).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":8081", router))
}
