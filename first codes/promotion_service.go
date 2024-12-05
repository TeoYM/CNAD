package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// PromotionService represents the promotion service
type PromotionService struct {
	db *sql.DB
}

// NewPromotionService returns a new promotion service
func NewPromotionService(db *sql.DB) *PromotionService {
	return &PromotionService{db: db}
}

// Promotion represents a promotion
type Promotion struct {
	ID          string    `json:"id"`
	Code        string    `json:"code"`
	Description string    `json:"description"`
	Discount    float64   `json:"discount"`
	ExpiryDate  time.Time `json:"expiry_date"`
	CreatedAt   time.Time `json:"created_at"`
}

// CreatePromotion handles promotion creation
func (s *PromotionService) CreatePromotion(w http.ResponseWriter, r *http.Request) {
	var promotion Promotion
	err := json.NewDecoder(r.Body).Decode(&promotion)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Save promotion to database
	_, err = s.db.Exec("INSERT INTO Promotions (Code, Description, Discount, ExpiryDate, CreatedAt) VALUES (?, ?, ?, ?, ?)", promotion.Code, promotion.Description, promotion.Discount, promotion.ExpiryDate, time.Now())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(promotion)
}

// GetPromotion handles promotion retrieval
func (s *PromotionService) GetPromotion(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	promotionID := params["id"]

	// Find promotion by ID
	var promotion Promotion
	var expiryDateStr string
	var createdAtStr string
	err := s.db.QueryRow("SELECT * FROM Promotions WHERE ID = ?", promotionID).Scan(&promotion.ID, &promotion.Code, &promotion.Description, &promotion.Discount, &expiryDateStr, &createdAtStr)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Promotion not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Parse expiry date string into time.Time object
	promotion.ExpiryDate, err = time.Parse("2006-01-02", expiryDateStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Parse created at date string into time.Time object
	promotion.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAtStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(promotion)
}

// UpdatePromotion handles promotion updates
func (s *PromotionService) UpdatePromotion(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	promotionID := params["id"]

	var promotion Promotion
	err := json.NewDecoder(r.Body).Decode(&promotion)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update promotion in database
	_, err = s.db.Exec("UPDATE Promotions SET Code = ?, Description = ?, Discount = ?, ExpiryDate = ? WHERE ID = ?", promotion.Code, promotion.Description, promotion.Discount, promotion.ExpiryDate, promotionID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(promotion)
}

// DeletePromotion handles promotion deletion
func (s *PromotionService) DeletePromotion(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	promotionID := params["id"]

	// Delete promotion from database
	_, err := s.db.Exec("DELETE FROM Promotions WHERE ID = ?", promotionID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func main() {
	// Connect to database
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/assg1")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create promotion service
	promotionService := NewPromotionService(db)

	// Create router
	router := mux.NewRouter()

	// Register routes
	router.HandleFunc("/promotions", promotionService.CreatePromotion).Methods("POST")
	router.HandleFunc("/promotions/{id}", promotionService.GetPromotion).Methods("GET")
	router.HandleFunc("/promotions/{id}", promotionService.UpdatePromotion).Methods("PUT")
	router.HandleFunc("/promotions/{id}", promotionService.DeletePromotion).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":8083", router))
}
