package models

import "gorm.io/gorm"

type TechProjects struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	LinkToGit   string `json:"link-to-git"`
}
