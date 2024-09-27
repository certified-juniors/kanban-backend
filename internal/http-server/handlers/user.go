package handlers

import (
	"encoding/json"
	"fmt"
	"kanban/internal/domain/models"
	"kanban/internal/lib/api/validation"
	"kanban/internal/lib/logger/sl"
	"net/http"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var op string = "User.Login"
	log := h.log.With("op", op)

	var credentials models.Credentials

	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		log.With("op", op).With("error", err).Error("error decoding body")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	authTokens, err := h.userService.Login(r.Context(), credentials)
	if err != nil {
		log.With("op", op).With("error", err).Error("error calling Login")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(authTokens)
}

// TODO: Add human-readable errors
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var op string = "User.Register"
	log := h.log.With("op", op)

	var credentials models.Credentials

	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if validationErrors := validation.ValidateStruct(credentials); len(validationErrors) > 0 {
		log.With("credentials", credentials).Error("invalid input parameters", sl.Err(fmt.Errorf("validation errors")))
		http.Error(w, "invalid input parameters", http.StatusBadRequest)
		return
	}

	authTokens, err := h.userService.Register(r.Context(), credentials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(authTokens)
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) Me(w http.ResponseWriter, r *http.Request) {

}
