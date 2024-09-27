package projectservice

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"kanban/internal/domain/models"
	"log/slog"
)

var (
	ErrFailedToUpdateProject = errors.New("failed to delete project")
)

func (s Service) Update(ctx context.Context, project models.Project, userId uuid.UUID) (*models.Project, error) {
	var op string = "ProjectService.Update"
	log := s.log.With("op", op)

	updatedProject, err := s.projectRepo.Update(ctx, project, userId)
	if err != nil {
		log.Error("error to update project", slog.String("error", err.Error()))
		return nil, ErrFailedToUpdateProject
	}

	return updatedProject, nil
}
