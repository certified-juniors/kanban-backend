package models

import "github.com/google/uuid"

type Project struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name" validate:"required"`
	Owner uuid.UUID `json:"owner" validate:"required"`
	Tasks []Task    `json:"tasks"`
}
