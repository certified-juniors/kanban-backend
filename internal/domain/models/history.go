package models

import "github.com/google/uuid"

type History struct {
	ID      uuid.UUID `json:"id"`
	Content string    `json:"content"`
	Author  uuid.UUID `json:"author"`
}
