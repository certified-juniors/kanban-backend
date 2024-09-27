package handlers

import (
	"encoding/json"
	"kanban/internal/domain/models"
	"net/http"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
}

// TODO: Add human-readable errors
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var credentials models.Credentials

	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	authTokens, err := h.userService.Register(r.Context(), credentials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(authTokens)
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) Me(w http.ResponseWriter, r *http.Request) {

}
