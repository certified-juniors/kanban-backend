package projectspostgresrepo

import (
	"context"
	"kanban/internal/domain/models"
)

func (r *ProjectsPostgresRepository) Insert(ctx context.Context, project models.Project) (*models.Project, error) {
	var createdProject models.Project
	var query string = `
		INSERT INTO projects
		(name, owner)
		VALUES ($1, $2)
		RETURNING id, name, owner
	`

	err := r.db.QueryRow(ctx, query, project.Name, project.Owner).Scan(&createdProject.ID, &createdProject.Name, &createdProject.Owner)
	if err != nil {
		return nil, err
	}

	return &createdProject, nil
}
