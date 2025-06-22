package database

import (
	"safehouse-main-back/src/internal/models"
)

type Database interface {
	GetContact() (*models.Contacts, error)
	GetTechProjects() ([]*models.ProjectGroups, error)
	GetFinanceProjects() ([]*models.ProjectGroups, error)
	GetGameProjects() ([]*models.ProjectGroups, error)
	GetGamesPlayed() ([]*models.GamesPlayed, error)
}
