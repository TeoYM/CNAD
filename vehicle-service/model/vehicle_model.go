package model

// Vehicle represents a vehicle
type Vehicle struct {
	ID           string `json:"id"`
	LicensePlate string `json:"license_plate"`
	VehicleType  string `json:"vehicle_type"`
	Availability bool   `json:"availability"`
	CreatedAt    string `json:"created_at"`
}
