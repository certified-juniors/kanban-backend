package projectspostgresrepo

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"kanban/internal/domain/models"
)

func (r *ProjectsPostgresRepository) Update(ctx context.Context, project models.Project, userId uuid.UUID) (*models.Project, error) {
	var selectRoleQuery string = `
		SELECT role FROM users_projects
		WHERE user_id = $1 AND project_id = $2
	`

	var role string

	err := r.db.QueryRow(ctx, selectRoleQuery, userId, project.ID).Scan(&role)
	if err != nil {
		return nil, err
	}

	if role != "CREATOR" {
		return nil, errors.New("invalid role")
	}

	var updateRowQuery string = `
		UPDATE projects
		SET name = $1
		WHERE id = $2
		RETURNING id, name
	`

	var updatedProject models.Project

	err = r.db.QueryRow(ctx, updateRowQuery, project.Name, project.ID).Scan(&updatedProject.ID, &updatedProject.Name)
	if err != nil {
		return nil, err
	}

	return &updatedProject, nil
}
