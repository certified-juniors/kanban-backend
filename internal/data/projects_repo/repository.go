package projectspostgresrepo

import (
	"context"
	"github.com/google/uuid"
	"kanban/internal/domain/models"
	"kanban/internal/lib/postgresql"
)

type ProjectsPostgresRepo interface {
	Insert(ctx context.Context, project models.Project) (*models.Project, error)
	Update(ctx context.Context, project models.Project) (*models.Project, error)
	Delete(ctx context.Context, id uuid.UUID) error
	GetById(ctx context.Context, id uuid.UUID) (*models.Project, error)
	GetAllBy(ctx context.Context) ([]models.Project, error)
}

type ProjectsPostgresRepository struct {
	db *postgresql.Postgres
}

func NewProjectsPostgresRepository(db *postgresql.Postgres) *ProjectsPostgresRepository {
	return &ProjectsPostgresRepository{db: db}
}
