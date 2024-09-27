package projectspostgresrepo

import (
	"context"
	"github.com/google/uuid"
	"kanban/internal/domain/models"
)

func (r *ProjectsPostgresRepository) Insert(ctx context.Context, project models.Project, creatorId uuid.UUID) (*models.Project, error) {
	var createdProject models.Project
	var insertProjectQuery string = `
		INSERT INTO projects
		(name)
		VALUES ($1)
		RETURNING id, name
	`

	err := r.db.QueryRow(ctx, insertProjectQuery, project.Name).Scan(&createdProject.ID, &createdProject.Name)
	if err != nil {
		return nil, err
	}

	var insertManyToManyQuery string = `
		INSERT INTO users_projects
		(user_id, project_id, role) 
		VALUES ($1, $2, $3)
	`

	_, err = r.db.Exec(ctx, insertManyToManyQuery, creatorId, createdProject.ID, "CREATOR")
	if err != nil {
		return nil, err
	}

	return &createdProject, nil
}
