package model

import (
	"time"
)

// Billing represents a billing record
type Billing struct {
	ID           string    `json:"id"`
	UserID       string    `json:"user_id"`
	VehicleID    string    `json:"vehicle_id"`
	RentalPeriod string    `json:"rental_period"`
	TotalCost    float64   `json:"total_cost"`
	CreatedAt    time.Time `json:"created_at"`
}
