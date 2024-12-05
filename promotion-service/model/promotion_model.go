package model

import (
	"time"
)

// Promotion represents a promotion
type Promotion struct {
	ID          string    `json:"id"`
	Code        string    `json:"code"`
	Description string    `json:"description"`
	Discount    float64   `json:"discount"`
	ExpiryDate  time.Time `json:"expiry_date"`
	CreatedAt   time.Time `json:"created_at"`
}
