package database

import (
	"safehouse-main-back/src/internal/models"
)

type Database interface {
	GetContact() (*models.Contacts, error)
	GetTechProjects() ([]*models.TechProjects, error)
	GetGames() ([]*models.Games, error)
}
