package taskspostgresrepo

import (
	"context"
	"kanban/internal/domain/models"
)

type TaskPostgresRepo interface {
	GetList(ctx context.Context, filters models.TaskFilters, offset int, limit int) (tasks []models.Task, totalCount int, err error)
}
