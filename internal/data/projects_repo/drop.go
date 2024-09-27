package projectspostgresrepo

import (
	"context"
	"errors"
	"github.com/google/uuid"
)

func (r *ProjectsPostgresRepository) Delete(ctx context.Context, projectId, ownerId uuid.UUID) error {
	var selectRoleQuery string = `
		SELECT role FROM users_projects
		WHERE user_id = $1 AND project_id = $2;
	`

	var role string

	err := r.db.QueryRow(ctx, selectRoleQuery, ownerId, projectId).Scan(&role)
	if err != nil {
		return err
	}

	if role != "CREATOR" {
		return nil
	}

	var dropRowQuery string = `
		DELETE FROM projects
		WHERE id = $1;
	`

	result, err := r.db.Exec(ctx, dropRowQuery, projectId)
	if err != nil {
		return err
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("no projects were deleted")
	}

	return nil
}
