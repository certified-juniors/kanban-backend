package projectservice

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"kanban/internal/lib/logger/sl"
	"log/slog"
)

var (
	ErrFailedToDeleteProject = errors.New("failed to delete project")
)

func (s Service) Delete(ctx context.Context, id, ownerId uuid.UUID) error {
	var op string = "ProjectService.Delete"
	log := s.log.With("op", op)

	err := s.projectRepo.Delete(ctx, id, ownerId)
	if err != nil {
		log.Error("error deleting project", slog.String("operation", op), sl.Err(err))
		return ErrFailedToDeleteProject
	}

	return nil
}
