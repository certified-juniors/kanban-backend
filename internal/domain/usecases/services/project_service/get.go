package projectservice

import (
	"context"
	"github.com/google/uuid"
	"kanban/internal/domain/models"
)

func (s Service) GetByID(ctx context.Context, id uuid.UUID) (*models.Project, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) GetAll(ctx context.Context) ([]*models.Project, error) {
	//TODO implement me
	panic("implement me")
}
