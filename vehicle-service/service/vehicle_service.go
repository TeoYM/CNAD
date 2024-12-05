package service

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"vehicle-service/model"
)

// VehicleService represents the vehicle service
type VehicleService struct {
	db *sql.DB
}

// NewVehicleService returns a new vehicle service
func NewVehicleService(db *sql.DB) *VehicleService {
	return &VehicleService{db: db}
}

// CreateVehicle creates a new vehicle
func (s *VehicleService) CreateVehicle(vehicle model.Vehicle) (model.Vehicle, error) {
	// Save vehicle to database
	_, err := s.db.Exec("INSERT INTO Vehicles (LicensePlate, VehicleType, Availability, CreatedAt) VALUES (?, ?, ?, ?)", vehicle.LicensePlate, vehicle.VehicleType, vehicle.Availability, time.Now())
	if err != nil {
		return vehicle, err
	}

	return vehicle, nil
}

// GetVehicle retrieves a vehicle by ID
func (s *VehicleService) GetVehicle(vehicleID string) (model.Vehicle, error) {
	log.Printf("Getting vehicle with ID: %s", vehicleID)

	// Find vehicle by ID
	var vehicle model.Vehicle
	var createdAt string
	err := s.db.QueryRow("SELECT * FROM Vehicles WHERE ID = ?", vehicleID).Scan(&vehicle.ID, &vehicle.LicensePlate, &vehicle.VehicleType, &vehicle.Availability, &createdAt)
	if err != nil {
		log.Printf("Error getting vehicle: %v", err)
		if err == sql.ErrNoRows {
			return vehicle, errors.New("vehicle not found")
		}
		return vehicle, err
	}

	vehicle.CreatedAt = createdAt
	log.Printf("Vehicle found: %+v", vehicle)

	return vehicle, nil
}

// UpdateVehicle updates a vehicle
func (s *VehicleService) UpdateVehicle(vehicleID string, vehicle model.Vehicle) (model.Vehicle, error) {
	// Update vehicle in database
	_, err := s.db.Exec("UPDATE Vehicles SET LicensePlate = ?, VehicleType = ?, Availability = ? WHERE ID = ?", vehicle.LicensePlate, vehicle.VehicleType, vehicle.Availability, vehicleID)
	if err != nil {
		return vehicle, err
	}

	return vehicle, nil
}

// DeleteVehicle deletes a vehicle by ID
func (s *VehicleService) DeleteVehicle(vehicleID string) error {
	// Delete vehicle from database
	_, err := s.db.Exec("DELETE FROM Vehicles WHERE ID = ?", vehicleID)
	if err != nil {
		return err
	}

	return nil
}
