package database

import (
	"errors"
	"gorm.io/gorm"
	"safehouse-main-back/src/internal/models"
)

type PostgresDB struct {
	db *gorm.DB
}

func NewPostgresDB(db *gorm.DB) Database {
	return &PostgresDB{db: db}
}

func (p *PostgresDB) GetContact() (*models.Contacts, error) {
	var contact models.Contacts
	if err := p.db.Where("active = ?", true).First(&contact).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		panic(err)
	}
	return &contact, nil
}

func (p *PostgresDB) GetProjects(projectType models.ProjectType) ([]*models.ProjectGroups, error) {
	var techProjectGroups []*models.ProjectGroups

	if err := p.db.
		Where("project_type = ?", projectType).
		Preload("TechRepositories").
		Order("created_at desc").
		Find(&techProjectGroups).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []*models.ProjectGroups{}, nil
		}
		panic(err)
	}

	return techProjectGroups, nil
}

func (p *PostgresDB) GetGamesPlayed() ([]*models.GamesPlayed, error) {
	var gamesPlayed []*models.GamesPlayed

	if err := p.db.
		Order("created_at desc").
		Limit(5). // limit for now, just in case
		Find(&gamesPlayed).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []*models.GamesPlayed{}, nil
		}
		panic(err)
	}

	return gamesPlayed, nil
}
