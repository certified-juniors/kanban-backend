package projectspostgresrepo

import (
	"context"
	"github.com/google/uuid"
	"kanban/internal/domain/models"
)

func (r *ProjectsPostgresRepository) GetAllBy(ctx context.Context) ([]models.Project, error) {
	return nil, nil
}

func (r *ProjectsPostgresRepository) GetById(ctx context.Context, uuid uuid.UUID) (*models.Project, error) {
	return nil, nil
}
