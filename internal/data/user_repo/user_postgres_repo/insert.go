package userpostgresrepo

import (
	"context"
	"github.com/google/uuid"
	"kanban/internal/domain/models"
)

func (r *UserPostgresRepository) Insert(ctx context.Context, credentials models.Credentials) (uuid.UUID, error) {
	return uuid.New(), nil
}
