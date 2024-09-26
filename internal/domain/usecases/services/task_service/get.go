package taskservice

import (
	"context"
	"kanban/internal/domain/models"
	"kanban/internal/lib/logger/sl"
	"log/slog"
)

func (s *Service) GetTaskList(ctx context.Context, filters models.TaskFilters, page int, limit int) (tasks []models.Task, totalPages int, err error) {
	const op = "Task.GetTaskList"
	log := s.log.With("op", op)

	offset := (page - 1) * limit

	tasks, totalCount, err := s.taskRepo.GetList(ctx, filters, offset, limit)
	if err != nil {
		log.Error("error get tasks list", slog.String("operation", op), sl.Err(err))
		return nil, 0, err
	}
	totalPages = (totalCount + limit - 1) / limit

	log.Info("tasks list received", slog.Int("totalPages", totalPages))

	return tasks, totalPages, nil
}
