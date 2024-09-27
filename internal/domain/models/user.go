package models

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name" validate:"required"`
	Surname    string    `json:"surname" validate:"required"`
	MiddleName string    `json:"middle_name" validate:"omitempty"`
	Email      string    `json:"email" validate:"required,email"`
	Password   string    `json:"password" validate:"required"`
	ProjectId  uuid.UUID `json:"project_id" validate:"omitempty"`
}

type Credentials struct {
	Name           string `json:"name" validate:"required"`
	Surname        string `json:"surname" validate:"required"`
	MiddleName     string `json:"middle_name" validate:"omitempty"`
	Email          string `json:"email" validate:"required,email"`
	Password       string `json:"password" validate:"required"`
	RepeatPassword string `json:"repeat_password" validate:"omitempty"`
}

type AuthTokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type JWTData struct {
	jwt.StandardClaims
	CustomClaims map[string]any `json:"custom_claims"`
}
