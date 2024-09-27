package userservice

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"kanban/internal/domain/models"
	"kanban/internal/lib/hash"
	customjwt "kanban/internal/lib/jwt"
	"kanban/internal/lib/logger/sl"
	"log/slog"
)

var (
	ErrAlreadyExists      = errors.New("user already exists")
	ErrFailedToCreateUser = errors.New("failed to create user")
	ErrUserNotFound       = errors.New("user not found")
)

// TODO: Add interaction with redis
// TODO: Add valiadtion for corner cases
func (s *Service) Register(ctx context.Context, credentials models.Credentials) (*models.AuthTokens, error) {
	var op string = "UserService.Register"
	log := s.log.With("op", op)

	candidate, err := s.userRepo.Exists(ctx, credentials.Email)
	if err != nil {
		log.Error("error to check user existence", slog.String("operation", op), sl.Err(err))
		return nil, ErrFailedToCreateUser
	}

	if credentials.Email == candidate.Email {
		log.Error("user already exists", slog.String("operation", op))
		return nil, ErrAlreadyExists
	}

	hashedPassword, err := hash.GenerateFromPassword(credentials.Password)
	if err != nil {
		log.Error("error to hash password", slog.String("operation", op), sl.Err(err))
		return nil, ErrFailedToCreateUser
	}

	credentials.Password = hashedPassword

	userID, err := s.userRepo.Insert(ctx, credentials)
	if err != nil {
		log.Error("error to insert user", slog.String("operation", op), sl.Err(err))
		return nil, ErrFailedToCreateUser
	}

	// TODO: Move Secret to config file
	authTokens, err := customjwt.GenerateJWTToken("SUPER_SECRET_KEY", userID)
	if err != nil {
		log.Error("error to create user token", slog.String("operation", op), sl.Err(err))
		return nil, ErrFailedToCreateUser
	}

	return authTokens, nil
}

// TODO: Add interaction with Redis
func (s *Service) Login(ctx context.Context, credentials models.Credentials) (*models.AuthTokens, error) {
	var op string = "UserService.Login"
	log := s.log.With("op", op)

	candidate, err := s.userRepo.Exists(ctx, credentials.Email)
	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		log.Error("error to check user existence", slog.String("operation", op), sl.Err(err))
		return nil, ErrUserNotFound
	}

	if err != nil {
		log.Error("error to check user existence", slog.String("operation", op), sl.Err(err))
		return nil, ErrFailedToCreateUser
	}

	if err = hash.CompareHashAndPassword(credentials.Password, credentials.Password); err != nil {
		log.Error("error to compare password", slog.String("operation", op), sl.Err(err))
		return nil, ErrFailedToCreateUser
	}

	// TODO: Move secret to config file
	authTokens, err := customjwt.GenerateJWTToken("SUPER_SECRET_KEY", candidate.ID)
	if err != nil {
		log.Error("error to create user token", slog.String("operation", op), sl.Err(err))
		return nil, ErrFailedToCreateUser
	}

	return authTokens, nil
}
