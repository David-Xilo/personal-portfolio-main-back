package models

import (
	"time"
)

type TechProjects struct {
	ID          uint       `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	LinkToGit   string     `json:"link-to-git"`
}
