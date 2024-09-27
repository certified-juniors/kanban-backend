package userservice

import (
	"context"
	userpostgresrepo "kanban/internal/data/user_repo/user_postgres_repo"
	"kanban/internal/domain/models"
	"log/slog"
)

type UserService interface {
	Register(ctx context.Context, credentials models.Credentials) (*models.AuthTokens, error)
}

type Service struct {
	log      *slog.Logger
	userRepo userpostgresrepo.UserPostgresRepo
}

func NewUserService(log *slog.Logger, userRepo userpostgresrepo.UserPostgresRepo) UserService {
	return &Service{
		log:      log,
		userRepo: userRepo,
	}
}
