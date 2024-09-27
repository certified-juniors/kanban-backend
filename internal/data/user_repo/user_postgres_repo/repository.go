package userpostgresrepo

import (
	"context"
	"github.com/google/uuid"
	"kanban/internal/domain/models"
	"kanban/internal/lib/postgresql"
)

type UserPostgresRepo interface {
	Insert(ctx context.Context, credentials models.Credentials) (uuid.UUID, error)
	Get(ctx context.Context, uuid uuid.UUID) (*models.User, error)
	Exists(ctx context.Context, email string) (*models.User, error)
}

type UserPostgresRepository struct {
	db *postgresql.Postgres
}

func NewUserPostgresRepository(db *postgresql.Postgres) *UserPostgresRepository {
	return &UserPostgresRepository{db: db}
}
