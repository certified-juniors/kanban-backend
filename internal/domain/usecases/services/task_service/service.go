package taskservice

import (
	"context"
	taskspostgresrepo "kanban/internal/data/task_repo/task_postgres_repo"
	"kanban/internal/domain/models"
	"log/slog"
)

type TaskService interface {
	GetTaskList(ctx context.Context, filters models.TaskFilters, page int, limit int) (staffs []models.Task, totalPages int, err error)
}

type Service struct {
	log      *slog.Logger
	taskRepo taskspostgresrepo.TaskPostgresRepo
}

func NewTaskService(log *slog.Logger, taskRepo taskspostgresrepo.TaskPostgresRepo) TaskService {
	return &Service{
		log:      log,
		taskRepo: taskRepo,
	}
}
