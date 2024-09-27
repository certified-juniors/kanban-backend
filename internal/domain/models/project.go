package models

import "github.com/google/uuid"

type Project struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name" validate:"required"`
	Owner uuid.UUID `json:"owner"`
	Tasks []Task    `json:"tasks"`
}
