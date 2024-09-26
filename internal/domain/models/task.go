package models

import (
	"github.com/gofrs/uuid"
	"time"
)

type Task struct {
	ID               uuid.UUID     `json:"id"`
	Name             string        `json:"name" required:"true"`
	Description      string        `json:"description,omitempty"`
	Attachments      []string      `json:"attachments,omitempty"`
	OriginalEstimate time.Duration `json:"original_estimate,omitempty"`
	TimeSpent        time.Duration `json:"time_spent,omitempty"`
	ParentID         uuid.UUID     `json:"parent_id,omitempty"`
	CreatedAt        time.Time     `json:"created_at,omitempty"`
	UpdatedAt        time.Time     `json:"updated_at,omitempty"`
	Author           uuid.UUID     `json:"author,omitempty"`
	Executor         uuid.UUID     `json:"executor,omitempty"`
	Status           TaskStatus    `json:"status,omitempty"`
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

type TaskStatus string

const (
	READY_FOR_ESTIMATE      TaskStatus = "READY_FOR_ESTIMATE"
	READY_FOR_DECOMPOSITION TaskStatus = "READY_FOR_DECOMPOSITION"
	READY_FOR_PROGRESS      TaskStatus = "READY_FOR_PROGRESS"
	IN_PROGRESS             TaskStatus = "IN_PROGRESS"
	COMPLETED               TaskStatus = "COMPLETED"
	READY_FOR_TEST          TaskStatus = "READY_FOR_TEST"
	IN_TESTING              TaskStatus = "IN_TESTING"
	READY_FOR_RELEASE       TaskStatus = "READY_FOR_RELEASE"
)
