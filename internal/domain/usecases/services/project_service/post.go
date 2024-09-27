package projectservice

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"kanban/internal/domain/models"
	"log/slog"
)

var (
	ErrFailedToCreate = errors.New("failed to create project")
)

func (s Service) Create(ctx context.Context, project models.Project, creatorId uuid.UUID) (*models.Project, error) {
	var op string = "ProjectsService.Create"
	log := s.log.With("op", op)

	createdProject, err := s.projectRepo.Insert(ctx, project, creatorId)
	if err != nil {
		log.Error("error to create project", slog.String("error", err.Error()))
		return nil, ErrFailedToCreate
	}

	return createdProject, nil
}
