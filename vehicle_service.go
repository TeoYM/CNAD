// vehicle_service.go

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// VehicleService represents the vehicle service
type VehicleService struct {
	// Add dependencies here (e.g., database connection)
}

// Vehicle represents a vehicle
type Vehicle struct {
	ID           string    `json:"id"`
	LicensePlate string    `json:"license_plate"`
	VehicleType  string    `json:"vehicle_type"`
	Availability bool      `json:"availability"`
	CreatedAt    time.Time `json:"created_at"`
}

// NewVehicleService returns a new vehicle service
func NewVehicleService() *VehicleService {
	return &VehicleService{}
}

// CreateVehicle handles vehicle creation
func (s *VehicleService) CreateVehicle(w http.ResponseWriter, r *http.Request) {
	var vehicle Vehicle
	err := json.NewDecoder(r.Body).Decode(&vehicle)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Save vehicle to database ( implementation omitted for brevity )
	// ...

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(vehicle)
}

// GetVehicle handles vehicle retrieval
func (s *VehicleService) GetVehicle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	vehicleID := params["id"]

	// Find vehicle by ID ( implementation omitted for brevity )
	// ...

	if vehicleID == "" {
		http.Error(w, "Vehicle not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Vehicle{ID: vehicleID})
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

	// Update vehicle in database ( implementation omitted for brevity )
	// ...

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vehicle)
}

// DeleteVehicle handles vehicle deletion
func (s *VehicleService) DeleteVehicle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	vehicleID := params["id"]

	// Delete vehicle from database ( implementation omitted for brevity )
	// ...

	w.WriteHeader(http.StatusNoContent)
}

func main() {
	router := mux.NewRouter()
	vehicleService := NewVehicleService()

	router.HandleFunc("/vehicles", vehicleService.CreateVehicle).Methods("POST")
	router.HandleFunc("/vehicles/{id}", vehicleService.GetVehicle).Methods("GET")
	router.HandleFunc("/vehicles/{id}", vehicleService.UpdateVehicle).Methods("PUT")
	router.HandleFunc("/vehicles/{id}", vehicleService.DeleteVehicle).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8081", router))
}
