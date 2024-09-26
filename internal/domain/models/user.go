package models

import "github.com/google/uuid"

type User struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name" validate:"required"`
	Surname    string    `json:"surname" validate:"required"`
	MiddleName string    `json:"middle_name" validate:"omitempty"`
	Email      string    `json:"email" validate:"required,email"`
	Password   string    `json:"password" validate:"required"`
	ProjectId  uuid.UUID `json:"project_id" validate:"omitempty"`
}
