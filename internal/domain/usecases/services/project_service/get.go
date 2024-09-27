package projectservice

import (
	"context"
	"github.com/google/uuid"
	"kanban/internal/domain/models"
	"log/slog"
)

func (s Service) GetByID(ctx context.Context, id uuid.UUID) (*models.Project, error) {
	var op string = "ProjectService.GetByID"
	log := s.log.With("op", op)

	project, err := s.projectRepo.GetById(ctx, id)
	if err != nil {
		log.Error("error while getting project by id", slog.String("error", err.Error()))
		return nil, err
	}

	return project, nil
}

func (s Service) GetAll(ctx context.Context, name string) ([]*models.Project, error) {
	var op string = "ProjectService.GetAll"
	log := s.log.With("op", op)

	projects, err := s.projectRepo.GetAllBy(ctx, name)
	if err != nil {
		log.Error("error while getting all projects", slog.String("error", err.Error()))
	}

	return projects, nil
}
