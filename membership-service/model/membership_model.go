package model

import (
	"time"
)

// Membership represents a membership
type Membership struct {
	ID             int       `json:"id"`
	UserID         int       `json:"user_id"`
	MembershipType string    `json:"membership_type"`
	CreatedAt      time.Time `json:"created_at"`
}
