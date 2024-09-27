package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"kanban/internal/domain/models"
	resp "kanban/internal/lib/api/response"
	"net/http"
)

// TODO: Add validation for fields

// CreateProject Создать новый проект
// @Summary Создание нового проекта
// @Description При создании нового проекта пользователь ставится в роль CREATOR этого проекта
// @Tags Projects
// @Router /projects [post]
// @Param project body models.CreateProjectsRequest true "Данные для создания нового проекта"
// @Success 201 {object} models.Project "Created"
// @Success 400 {object} resp.Response "Bad request"
// @Success 500 {object} resp.Response "Internal server error"
func (h *Handler) CreateProject(w http.ResponseWriter, r *http.Request) {
	var op string = "Projects.CreateProject"
	log := h.log.With("op", op)

	creatorId, ok := r.Context().Value("id").(uuid.UUID)
	if !ok {
		log.With("op", op).Error("missing id in context while creating project")
		resp.WriteJSONResponse(w, http.StatusUnauthorized, "user not authorized", nil)
		return
	}

	if creatorId == uuid.Nil {
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

	addedProject, err := h.projectsService.Create(r.Context(), project, creatorId)
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

func (h *Handler) UpdateProject(w http.ResponseWriter, r *http.Request) {
	var op string = "Projects.UpdateProject"
	log := h.log.With("op", op)

	userId, ok := r.Context().Value("id").(uuid.UUID)
	if !ok {
		log.With("op", op).Error("missing id in context while updating project")
		resp.WriteJSONResponse(w, http.StatusUnauthorized, "user not authorized", nil)
		return
	}

	if userId == uuid.Nil {
		log.With("op", op).Error("missing id in context while updating project")
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

	updatedProject, err := h.projectsService.Update(r.Context(), project, userId)
	if err != nil {
		log.With("op", op).With("error", err).Error("error updating project")
		resp.WriteJSONResponse(w, http.StatusInternalServerError, "error updating project", nil)
		return
	}

	resp.WriteJSON(w, http.StatusOK, updatedProject)
}

func (h *Handler) DeleteProject(w http.ResponseWriter, r *http.Request) {
	var op string = "Projects.DeleteProject"
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

	idParam := chi.URLParam(r, "id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		log.With("op", op).With("error", err).Error("error parsing id")
		resp.WriteJSONResponse(w, http.StatusBadRequest, "invalid id parameter", nil)
		return
	}

	err = h.projectsService.Delete(r.Context(), id, ownerId)
	if err != nil {
		log.With("op", op).With("error", err).Error("error deleting project")
		resp.WriteJSONResponse(w, http.StatusInternalServerError, "error deleting project", nil)
		return
	}

	resp.WriteJSONResponse(w, http.StatusOK, "project deleted", nil)
}
