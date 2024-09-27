package handlers

import (
	projectservice "kanban/internal/domain/usecases/services/project_service"
	taskservice "kanban/internal/domain/usecases/services/task_service"
	userservice "kanban/internal/domain/usecases/services/user_service"
	"log/slog"
)

type Handler struct {
	log             *slog.Logger
	taskService     taskservice.TaskService
	userService     userservice.UserService
	projectsService projectservice.ProjectService
}

func NewHandler(
	log *slog.Logger,
	taskService taskservice.TaskService,
	userService userservice.UserService,
	projectsService projectservice.ProjectService,
) *Handler {
	return &Handler{
		log:             log,
		taskService:     taskService,
		userService:     userService,
		projectsService: projectsService,
	}
}
