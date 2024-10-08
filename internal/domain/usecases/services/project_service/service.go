package projectservice

import (
	"context"
	"github.com/google/uuid"
	projectspostgresrepo "kanban/internal/data/projects_repo"
	"kanban/internal/domain/models"
	"log/slog"
)

type ProjectService interface {
	Create(ctx context.Context, project models.Project, creatorId uuid.UUID) (*models.Project, error)
	Update(ctx context.Context, project models.Project, userId uuid.UUID) (*models.Project, error)
	GetByID(ctx context.Context, id uuid.UUID) (*models.Project, error)
	GetAll(ctx context.Context, name string) ([]*models.Project, error)
	Delete(ctx context.Context, id, owner uuid.UUID) error
}

type Service struct {
	log         *slog.Logger
	projectRepo projectspostgresrepo.ProjectsPostgresRepo
}

func NewProjectsService(log *slog.Logger, projectRepo projectspostgresrepo.ProjectsPostgresRepo) *Service {
	return &Service{
		log:         log,
		projectRepo: projectRepo,
	}
}
