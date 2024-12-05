package service

import (
	"database/sql"
	"errors"
	"time"

	"membership-service/model"
)

// MembershipService represents the membership service
type MembershipService struct {
	db *sql.DB
}

// NewMembershipService returns a new membership service
func NewMembershipService(db *sql.DB) *MembershipService {
	return &MembershipService{db: db}
}

// CreateMembership creates a new membership
func (s *MembershipService) CreateMembership(membership model.Membership) error {
	stmt, err := s.db.Prepare("INSERT INTO Memberships (UserID, MembershipType) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(membership.UserID, membership.MembershipType)
	if err != nil {
		return err
	}

	return nil
}

// GetMembership retrieves a membership by user ID
func (s *MembershipService) GetMembership(userID string) (model.Membership, error) {
	var membership model.Membership
	var createdAtStr string
	stmt, err := s.db.Prepare("SELECT * FROM Memberships WHERE UserID = ?")
	if err != nil {
		return membership, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(userID).Scan(&membership.ID, &membership.UserID, &membership.MembershipType, &createdAtStr)
	if err != nil {
		if err == sql.ErrNoRows {
			return membership, errors.New("membership not found")
		}
		return membership, err
	}

	membership.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAtStr)
	if err != nil {
		return membership, err
	}

	return membership, nil
}
