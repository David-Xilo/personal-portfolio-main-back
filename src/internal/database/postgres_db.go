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

func (p *PostgresDB) GetGameProjects() ([]*models.ProjectGroups, error) {
	var gameProjectGroups []*models.ProjectGroups

	if err := p.db.
		Preload("GameRepositories").
		Joins("JOIN game_projects ON project_groups.id = game_projects.project_group_id").
		Order("created_at desc").
		//Limit(10).
		Find(&gameProjectGroups).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []*models.ProjectGroups{}, nil
		}
		panic(err)
	}

	return gameProjectGroups, nil
}

func (p *PostgresDB) GetTechProjects() ([]*models.ProjectGroups, error) {
	var techProjectGroups []*models.ProjectGroups

	if err := p.db.
		Where("project_type = ?", "tech").
		Preload("TechProjects").
		// REMOVE THIS LINE: Joins("JOIN tech_projects ON project_groups.id = tech_projects.project_group_id").
		Order("created_at desc").
		Find(&techProjectGroups).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []*models.ProjectGroups{}, nil
		}
		panic(err)
	}

	return techProjectGroups, nil
}

func (p *PostgresDB) GetFinanceProjects() ([]*models.ProjectGroups, error) {
	var financeProjectGroups []*models.ProjectGroups

	if err := p.db.
		Preload("FinanceRepositories").
		Joins("JOIN finance_projects ON project_groups.id = finance_projects.project_group_id").
		Order("created_at desc").
		//Limit(10).
		Find(&financeProjectGroups).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []*models.ProjectGroups{}, nil
		}
		panic(err)
	}

	return financeProjectGroups, nil
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
