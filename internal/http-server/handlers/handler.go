package handlers

import (
	taskservice "kanban/internal/domain/usecases/services/task_service"
	userservice "kanban/internal/domain/usecases/services/user_service"
	"log/slog"
)

type Handler struct {
	log         *slog.Logger
	taskService taskservice.TaskService
	userService userservice.UserService
}

func NewHandler(
	log *slog.Logger,
	taskService taskservice.TaskService,
	userService userservice.UserService,
) *Handler {
	return &Handler{
		log:         log,
		taskService: taskService,
		userService: userService,
	}
}
