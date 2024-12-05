package service

import (
	"database/sql"
	"errors"
	"time"

	"billing-service/model"
)

// BillingService represents the billing service
type BillingService struct {
	db *sql.DB
}

// NewBillingService returns a new billing service
func NewBillingService(db *sql.DB) *BillingService {
	return &BillingService{db: db}
}

// CreateBillingRecord creates a new billing record
func (s *BillingService) CreateBillingRecord(billing model.Billing) error {
	billing.CreatedAt = time.Now() // Set the CreatedAt field to the current timestamp

	stmt, err := s.db.Prepare("INSERT INTO Billings (UserID, VehicleID, RentalPeriod, TotalCost, CreatedAt) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(billing.UserID, billing.VehicleID, billing.RentalPeriod, billing.TotalCost, billing.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

// GetBillingRecord retrieves a billing record by ID
func (s *BillingService) GetBillingRecord(billingID string) (model.Billing, error) {
	var billing model.Billing
	var createdAt string

	stmt, err := s.db.Prepare("SELECT * FROM Billings WHERE ID = ?")
	if err != nil {
		return billing, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(billingID).Scan(&billing.ID, &billing.UserID, &billing.VehicleID, &billing.RentalPeriod, &billing.TotalCost, &createdAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return billing, errors.New("billing record not found")
		}
		return billing, err
	}

	t, err := time.Parse("2006-01-02 15:04:05", createdAt)
	if err != nil {
		return billing, err
	}

	billing.CreatedAt = t

	return billing, nil
}
