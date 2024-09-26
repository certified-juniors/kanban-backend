package handlers

import (
	taskservice "kanban/internal/domain/usecases/services/task_service"
	"log/slog"
)

type Handler struct {
	log         *slog.Logger
	taskService taskservice.TaskService
}

func NewHandler(
	log *slog.Logger,
	taskService taskservice.TaskService,
) *Handler {
	return &Handler{
		log:         log,
		taskService: taskService,
	}
}
