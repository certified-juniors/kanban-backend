package taskspostgresrepo

import (
	"context"
	"github.com/pkg/errors"
	"kanban/internal/domain/models"
	"kanban/internal/lib/postgresql"
)

var (
	ErrTaskNotFound = errors.New("task not found")
)

type TaskPostgresRepository struct {
	db *postgresql.Postgres
}

func NewTaskPostgresRepository(db *postgresql.Postgres) *TaskPostgresRepository {
	return &TaskPostgresRepository{
		db: db,
	}
}

func (r *TaskPostgresRepository) GetList(ctx context.Context, filters models.TaskFilters, offset int, limit int) (tasks []models.Task, totalCount int, err error) {
	//var query string
	//var rows *sql.Rows
	//
	//query = `
	//	WITH StaffList AS (
	//		SELECT
	//			u.Id, u.Name, u.Surname, u.Login,
	//			u.Password, u.shopcode, u.Status,
	//			u.Role, u.userINN, u.Patronymic,
	//			u.brandId,
	//			COUNT(*) OVER() AS total_count
	//		FROM Users u
	//		WHERE 1 = 1
	//`
	//
	//args := []interface{}{}
	//
	//if filters.BrandID != nil {
	//	query += ` AND u.brandId = ?`
	//	args = append(args, *filters.BrandID)
	//}
	//
	//if filters.Login != nil {
	//	query += ` AND UPPER(u.Login) LIKE UPPER(?)`
	//	args = append(args, *filters.Login+"%")
	//}
	//
	//if filters.Status != nil {
	//	query += ` AND u.Status = ?`
	//	args = append(args, *filters.Status)
	//}
	//
	//if filters.Role != nil {
	//	query += ` AND UPPER(u.Role) LIKE UPPER(?)`
	//	args = append(args, *filters.Role+"%")
	//}
	//
	//if filters.ShopCode != nil {
	//	query += ` AND UPPER(u.shopcode) LIKE UPPER(?)`
	//	args = append(args, *filters.ShopCode+"%")
	//}
	//
	//query += `
	//	)
	//	SELECT
	//		u.Id, u.Name, u.Surname, u.Login,
	//		u.Password, u.shopcode, u.Status,
	//		u.Role, u.userINN, u.Patronymic, u.brandId, u.total_count
	//	FROM StaffList u
	//	ORDER BY u.Id DESC
	//	OFFSET ? ROWS FETCH NEXT ? ROWS ONLY
	//`
	//
	//args = append(args, offset, limit)
	//
	//rows, err = r.mssql.QueryContext(ctx, query, args...)
	//if err != nil {
	//	if errors.Is(err, sql.ErrNoRows) {
	//		return nil, 0, nil
	//	}
	//	return nil, 0, errors.New("error querying staff list: " + err.Error())
	//}
	//defer rows.Close()
	//
	//var totalCountFromQuery int
	//
	//for rows.Next() {
	//	var staff entities.Staff
	//	if err = rows.Scan(
	//		&staff.ID,
	//		&staff.Name,
	//		&staff.Surname,
	//		&staff.Login,
	//		&staff.Password,
	//		&staff.ShopCode,
	//		&staff.Status,
	//		&staff.Role,
	//		&staff.UserINN,
	//		&staff.Patronymic,
	//		&staff.BrandID,
	//		&totalCountFromQuery,
	//	); err != nil {
	//		return nil, 0, errors.New("error scanning row: " + err.Error())
	//	}
	//	staffList = append(staffList, staff)
	//}
	//
	//if len(staffList) > 0 {
	//	totalCount = totalCountFromQuery
	//}
	//
	//if err = rows.Err(); err != nil {
	//	return nil, 0, errors.New("error iterating over rows: " + err.Error())
	//}
	//
	//return staffList, totalCount, nil

	return nil, 0, nil
}
