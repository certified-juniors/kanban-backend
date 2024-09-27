package handlers

import (
	"encoding/json"
	"github.com/google/uuid"
	"kanban/internal/domain/models"
	resp "kanban/internal/lib/api/response"
	"net/http"
)

// TODO: Add validation for fields
func (h *Handler) CreateProject(w http.ResponseWriter, r *http.Request) {
	var op string = "Projects.CreateProject"
	log := h.log.With("op", op)

	ownerId, ok := r.Context().Value("id").(uuid.UUID)
	if !ok {
		log.With("op", op).Error("missing id in context while creating project")
		resp.WriteJSONResponse(w, http.StatusUnauthorized, "user not authorized", nil)
		return
	}

	if ownerId == uuid.Nil {
		log.With("op", op).Error("missing id in context while creating project")
		resp.WriteJSONResponse(w, http.StatusUnauthorized, "user not authorized", nil)
		return
	}

	var project models.Project

	err := json.NewDecoder(r.Body).Decode(&project)
	if err != nil {
		log.With("op", op).With("error", err).Error("error decoding body")
		resp.WriteJSONResponse(w, http.StatusBadRequest, "invalid input parameters", nil)
		return
	}

	project.Owner = ownerId

	addedProject, err := h.projectsService.Create(r.Context(), project)
	if err != nil {
		log.With("op", op).With("error", err).Error("error creating project")
		resp.WriteJSONResponse(w, http.StatusInternalServerError, "error creating project", nil)
		return
	}

	resp.WriteJSON(w, http.StatusCreated, addedProject)
}

func (h *Handler) GetProjects(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) GetProject(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) UpdateProject(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) DeleteProject(w http.ResponseWriter, r *http.Request) {}
