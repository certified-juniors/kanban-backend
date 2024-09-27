package userservice

import (
	"context"
	userpostgresrepo "kanban/internal/data/user_repo/user_postgres_repo"
	"kanban/internal/domain/models"
	"log/slog"
)

type UserService interface {
	Register(ctx context.Context, credentials models.Credentials) (*models.AuthTokens, error)
	Login(ctx context.Context, credentials models.Credentials) (*models.AuthTokens, error)
	Logout(ctx context.Context) error
	Me(ctx context.Context) (*models.User, error)
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
