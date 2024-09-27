package models

type LoginRequest struct {
	Email    string `json:"email" example:"user@example.com"`
	Password string `json:"password" example:"password123"`
}

type RegisterRequest struct {
	Email          string `json:"email" example:"user@example.com"`
	Password       string `json:"password" example:"password123"`
	Name           string `json:"name" example:"user"`
	Surname        string `json:"surname" example:"user"`
	MiddleName     string `json:"middle_name" example:"user"`
	RepeatPassword string `json:"password" example:"password123"`
}

type CreateProjectsRequest struct {
	Name string `json:"name" example:"New project"`
}
