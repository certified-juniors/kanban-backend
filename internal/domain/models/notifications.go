package models

import "github.com/google/uuid"

type Notification struct {
	ID    uuid.UUID `json:"id"`
	Title string    `json:"title" required:"true"`
}
