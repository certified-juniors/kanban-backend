package userpostgresrepo

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"kanban/internal/domain/models"
)

func (r *UserPostgresRepository) Get(ctx context.Context, uuid uuid.UUID) (*models.User, error) {
	return &models.User{}, nil
}

func (r *UserPostgresRepository) Exists(ctx context.Context, email string) (*models.User, error) {
	var query string
	var rows pgx.Rows
	var user models.User

	query = `
	SELECT id, name, surname, middle_name, email, password FROM users
	WHERE email = $1
	`

	rows, err := r.db.Query(ctx, query, email)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Surname, &user.MiddleName, &user.Email, &user.Password); err != nil {
			return nil, err
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &user, nil
}
