package projectspostgresrepo

import (
	"context"
	"github.com/google/uuid"
	"kanban/internal/domain/models"
)

func (r *ProjectsPostgresRepository) GetAllBy(ctx context.Context, name string) ([]*models.Project, error) {
	var projects []*models.Project
	var query string = `
	SELECT * FROM projects
	WHERE name = $1
	`

	rows, err := r.db.Query(ctx, query, name)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var project *models.Project
		if err := rows.Scan(&project.ID, name); err != nil {
			return nil, err
		}

		projects = append(projects, project)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return projects, nil
}

func (r *ProjectsPostgresRepository) GetById(ctx context.Context, uuid uuid.UUID) (*models.Project, error) {
	var project models.Project
	var query string = `
	SELECT * FROM projects
	WHERE id = $1
	`

	rows, err := r.db.Query(ctx, query, uuid)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&project.ID, uuid); err != nil {
			return nil, err
		}
	}

	return &project, nil
}
