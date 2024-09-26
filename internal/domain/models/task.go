package models

type Task struct {
}

type TaskFilters struct {
	TaskCode *string `json:"task_code" validate:"omitempty"`
	Name     *string `json:"name" validate:"omitempty"`
	UserID   *int    `json:"user_id" validate:"omitempty"`
}

type TaskPagination struct {
	TotalPages int    `json:"total_pages"`
	Tasks      []Task `json:"tasks"`
}
