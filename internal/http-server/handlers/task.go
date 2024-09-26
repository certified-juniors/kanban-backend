package handlers

import (
	"fmt"
	"kanban/internal/domain/models"
	resp "kanban/internal/lib/api/response"
	apiutils "kanban/internal/lib/api/utils"
	"kanban/internal/lib/api/validation"
	"kanban/internal/lib/logger/sl"
	"net/http"
)

// GetTaskList Получить список задач
// @Summary Получить список задач
// @Description Получить список задач с возможностью фильтрациями
// @Tags Task
// @Router /task/getList [get]
// @Param page query int false "Номер страницы (default: 1)"
// @Param limit query int false "Количество элементов на странице (default: 20)"
// @Param task_code query string false "Код задача"
// @Param name query string false "Наименование"
// @Param category query string false "Категория"
// @Param user_id query int false "ID пользователя"
// @Success 200 {object} models.TaskPagination "OK"
// @Failure 400 {object} resp.Response "Bad Request"
// @Failure 500 {object} resp.Response "Internal Server Error"
func (h *Handler) GetTaskList(w http.ResponseWriter, r *http.Request) {
	const op = "Handler.GetTaskList"
	log := h.log.With("op", op)

	var filters models.TaskFilters
	query := r.URL.Query()

	if taskCode := query.Get("task_code"); taskCode != "" {
		filters.TaskCode = &taskCode
	}
	if name := query.Get("name"); name != "" {
		filters.Name = &name
	}

	apiutils.ParseIntParam(query.Get("user_id"), &filters.UserID, "user_id", log, w)
	page, limit := apiutils.ParsePagination(query.Get("page"), query.Get("limit"))

	if validationErrors := validation.ValidateStruct(filters); len(validationErrors) > 0 {
		log.With("filters", filters).Error("invalid input parameters", sl.Err(fmt.Errorf("validation errors")))
		resp.WriteJSONResponse(w, http.StatusBadRequest, "invalid input parameters", validationErrors)
		return
	}

	tasks, totalPages, err := h.taskService.GetTaskList(r.Context(), filters, page, limit)
	if err != nil {
		log.With("filters", filters).Error("Error getting task list", sl.Err(err))
		resp.WriteJSONResponse(w, http.StatusInternalServerError, "Error getting task list", nil)
		return
	}

	resp.WriteJSON(w, http.StatusOK, models.TaskPagination{
		TotalPages: totalPages,
		Tasks:      tasks,
	})
}
