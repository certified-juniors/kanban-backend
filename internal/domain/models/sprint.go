package models

import (
	"github.com/google/uuid"
	"time"
)

type Sprint struct {
	ID       uuid.UUID     `json:"id"`
	Name     string        `json:"name" required:"true"`
	Tasks    []Task        `json:"tasks"`
	Duration time.Duration `json:"duration" required:"true"`
	Status   SprintStatus  `json:"status"`
}

type SprintStatus string

const (
	InProgress SprintStatus = "in-progress"
	Completed  SprintStatus = "completed"
	Closed     SprintStatus = "closed"
)
