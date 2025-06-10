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

func (p *PostgresDB) GetGames() ([]*models.Games, error) {
	var games []*models.Games

	if err := p.db.
		Order("created_at desc").
		Limit(5).
		Find(&games).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []*models.Games{}, nil
		}
		panic(err)
	}

	return games, nil
}

func (p *PostgresDB) GetTechProjects() ([]*models.TechProjects, error) {
	var projects []*models.TechProjects

	if err := p.db.Limit(10).Find(&projects).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []*models.TechProjects{}, nil
		}
		panic(err)
	}

	return projects, nil
}

func (p *PostgresDB) GetFinanceProjects() ([]*models.FinanceProjects, error) {
	var projects []*models.FinanceProjects

	if err := p.db.Limit(10).Find(&projects).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []*models.FinanceProjects{}, nil
		}
		panic(err)
	}

	return projects, nil
}
