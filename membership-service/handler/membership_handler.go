package handler

import (
	"encoding/json"
	"net/http"

	"membership-service/model"
	"membership-service/service"

	"github.com/gorilla/mux"
)

// MembershipHandler represents the membership handler
type MembershipHandler struct {
	membershipService *service.MembershipService
}

// NewMembershipHandler returns a new membership handler
func NewMembershipHandler(membershipService *service.MembershipService) *MembershipHandler {
	return &MembershipHandler{membershipService: membershipService}
}

// CreateMembership handles membership creation
func (h *MembershipHandler) CreateMembership(w http.ResponseWriter, r *http.Request) {
	var membership model.Membership
	err := json.NewDecoder(r.Body).Decode(&membership)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.membershipService.CreateMembership(membership)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(membership)
}

// GetMembership handles membership retrieval
func (h *MembershipHandler) GetMembership(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["user_id"]

	membership, err := h.membershipService.GetMembership(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(membership)
}
