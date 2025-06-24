package models

import (
	"time"
)

type TechRepositories struct {
	ID             uint       `json:"id"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at,omitempty"`
	ProjectGroupID uint       `json:"project_group_id"`
	Title          string     `json:"title"`
	Description    string     `json:"description"`
	LinkToGit      string     `json:"link_to_git"`
}

type TechRepositoriesDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	LinkToGit   string `json:"link_to_git"`
}
