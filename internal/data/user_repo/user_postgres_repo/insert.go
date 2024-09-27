package userpostgresrepo

import (
	"context"
	"github.com/google/uuid"
	"kanban/internal/domain/models"
)

func (r *UserPostgresRepository) Insert(ctx context.Context, credentials models.Credentials) (uuid.UUID, error) {
	var id uuid.UUID
	var query string = `
	INSERT INTO users
	(name, surname, middle_name, email, password)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id;
	`

	err := r.db.QueryRow(ctx, query,
		credentials.Name,
		credentials.Surname,
		credentials.MiddleName,
		credentials.Email,
		credentials.Password,
	).Scan(&id)

	if err != nil {
		return id, err
	}

	return id, nil
}
