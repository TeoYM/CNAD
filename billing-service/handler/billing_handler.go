package handler

import (
	"encoding/json"
	"net/http"

	"billing-service/model"
	"billing-service/service"

	"github.com/gorilla/mux"
)

// BillingHandler represents the billing handler
type BillingHandler struct {
	billingService *service.BillingService
}

// NewBillingHandler returns a new billing handler
func NewBillingHandler(billingService *service.BillingService) *BillingHandler {
	return &BillingHandler{billingService: billingService}
}

// CreateBillingRecord handles billing record creation
func (h *BillingHandler) CreateBillingRecord(w http.ResponseWriter, r *http.Request) {
	var billing model.Billing
	err := json.NewDecoder(r.Body).Decode(&billing)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.billingService.CreateBillingRecord(billing)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(billing)
}

// GetBillingRecord handles billing record retrieval
func (h *BillingHandler) GetBillingRecord(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	billingID := params["id"]

	billing, err := h.billingService.GetBillingRecord(billingID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(billing)
}
