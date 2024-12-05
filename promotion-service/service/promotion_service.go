package service

import (
	"database/sql"
	"errors"
	"time"

	"promotion-service/model"
)

// PromotionService represents the promotion service
type PromotionService struct {
	db *sql.DB
}

// NewPromotionService returns a new promotion service
func NewPromotionService(db *sql.DB) *PromotionService {
	return &PromotionService{db: db}
}

// CreatePromotion creates a new promotion
func (s *PromotionService) CreatePromotion(promotion model.Promotion) error {
	if promotion.ExpiryDate.IsZero() {
		promotion.ExpiryDate = time.Now().AddDate(0, 0, 1) // set expiry date to tomorrow
	}

	stmt, err := s.db.Prepare("INSERT INTO Promotions (Code, Description, Discount, ExpiryDate, CreatedAt) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(promotion.Code, promotion.Description, promotion.Discount, promotion.ExpiryDate, time.Now())
	if err != nil {
		return err
	}

	return nil
}

// GetPromotion retrieves a promotion by ID
func (s *PromotionService) GetPromotion(promotionID string) (model.Promotion, error) {
	var promotion model.Promotion
	var expiryDateStr string
	var createdAtStr string
	stmt, err := s.db.Prepare("SELECT * FROM Promotions WHERE ID = ?")
	if err != nil {
		return promotion, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(promotionID).Scan(&promotion.ID, &promotion.Code, &promotion.Description, &promotion.Discount, &expiryDateStr, &createdAtStr)
	if err != nil {
		if err == sql.ErrNoRows {
			return promotion, errors.New("promotion not found")
		}
		return promotion, err
	}

	promotion.ExpiryDate, err = time.Parse("2006-01-02", expiryDateStr)
	if err != nil {
		return promotion, err
	}

	promotion.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAtStr)
	if err != nil {
		return promotion, err
	}

	return promotion, nil
}

// UpdatePromotion updates a promotion
func (s *PromotionService) UpdatePromotion(promotionID string, promotion model.Promotion) error {
	stmt, err := s.db.Prepare("UPDATE Promotions SET Code = ?, Description = ?, Discount = ?, ExpiryDate = ? WHERE ID = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(promotion.Code, promotion.Description, promotion.Discount, promotion.ExpiryDate, promotionID)
	if err != nil {
		return err
	}

	return nil
}

// DeletePromotion deletes a promotion
func (s *PromotionService) DeletePromotion(promotionID string) error {
	stmt, err := s.db.Prepare("DELETE FROM Promotions WHERE ID = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(promotionID)
	if err != nil {
		return err
	}

	return nil
}
