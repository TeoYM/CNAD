package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"vehicle-service/model"
	"vehicle-service/service"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var vehicleService *service.VehicleService

func init() {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/assg1")
	if err != nil {
		log.Fatal(err)
	}
	vehicleService = service.NewVehicleService(db)
}

// CreateVehicle handles vehicle creation
func CreateVehicle(w http.ResponseWriter, r *http.Request) {
	var vehicle model.Vehicle
	err := json.NewDecoder(r.Body).Decode(&vehicle)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdVehicle, err := vehicleService.CreateVehicle(vehicle)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdVehicle)
}

// GetVehicle handles vehicle retrieval
func GetVehicle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	vehicleID := params["id"]

	vehicle, err := vehicleService.GetVehicle(vehicleID)
	if err != nil {
		http.Error(w, "Vehicle not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vehicle)
}

// UpdateVehicle handles vehicle updates
func UpdateVehicle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	vehicleID := params["id"]

	var vehicle model.Vehicle
	err := json.NewDecoder(r.Body).Decode(&vehicle)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedVehicle, err := vehicleService.UpdateVehicle(vehicleID, vehicle)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedVehicle)
}

// DeleteVehicle handles vehicle deletion
func DeleteVehicle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	vehicleID := params["id"]

	err := vehicleService.DeleteVehicle(vehicleID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
