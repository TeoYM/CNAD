package handler

import (
	"encoding/json"
	"net/http"

	"promotion-service/model"
	"promotion-service/service"

	"github.com/gorilla/mux"
)

// PromotionHandler represents the promotion handler
type PromotionHandler struct {
	promotionService *service.PromotionService
}

// NewPromotionHandler returns a new promotion handler
func NewPromotionHandler(promotionService *service.PromotionService) *PromotionHandler {
	return &PromotionHandler{promotionService: promotionService}
}

// CreatePromotion handles promotion creation
func (h *PromotionHandler) CreatePromotion(w http.ResponseWriter, r *http.Request) {
	var promotion model.Promotion
	err := json.NewDecoder(r.Body).Decode(&promotion)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.promotionService.CreatePromotion(promotion)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(promotion)
}

// GetPromotion handles promotion retrieval
func (h *PromotionHandler) GetPromotion(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	promotionID := params["id"]

	promotion, err := h.promotionService.GetPromotion(promotionID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(promotion)
}

// UpdatePromotion handles promotion updates
func (h *PromotionHandler) UpdatePromotion(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	promotionID := params["id"]

	var promotion model.Promotion
	err := json.NewDecoder(r.Body).Decode(&promotion)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.promotionService.UpdatePromotion(promotionID, promotion)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(promotion)
}

// DeletePromotion handles promotion deletion
func (h *PromotionHandler) DeletePromotion(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	promotionID := params["id"]

	err := h.promotionService.DeletePromotion(promotionID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
