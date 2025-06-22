package mockdb

import (
	"safehouse-main-back/src/internal/database"
	"safehouse-main-back/src/internal/models"
	"time"
)

type MockDB struct{}

func NewMockDB() database.Database {
	return &MockDB{}
}

func (m *MockDB) GetContact() (*models.Contacts, error) {
	return &models.Contacts{
		ID:        1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      "Mock user",
		Email:     "mock@example.com",
		LinkedIn:  "linkedin.com/mockuser",
		Github:    "github.com/mockuser",
	}, nil
}

func (m *MockDB) GetTechProjects() ([]*models.ProjectGroups, error) {
	return []*models.ProjectGroups{
		{
			ID:          1,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			DeletedAt:   nil, // Indicates the project is not deleted
			Title:       "Sample Project",
			Description: "This is a sample project used for testing purposes.",
			//LinkToGit:   "https://github.com/sample/sample-project",
		},
	}, nil
}

func (m *MockDB) GetFinanceProjects() ([]*models.ProjectGroups, error) {
	return []*models.ProjectGroups{
		{},
	}, nil
}

func (m *MockDB) GetGameProjects() ([]*models.ProjectGroups, error) {
	return []*models.ProjectGroups{
		{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: nil, // Indicates the game is not deleted
			Title:     "Epic Adventure",
			//Genre:       models.GameGenreStrategy,
			Description: "Embark on an epic journey through uncharted lands.",
			//LinkToGit:   "https://github.com/example/epic-adventure",
			//LinkToStore: "https://store.example.com/epic-adventure",
		},
		{
			ID:        2,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: nil,
			Title:     "Mystic Quest",
			//Genre:       models.GameGenreTableTop,
			Description: "Dive into the mystic world and unravel its secrets.",
			//LinkToGit:   "https://github.com/example/mystic-quest",
			//LinkToStore: "https://store.example.com/mystic-quest",
		},
	}, nil
}

func (m *MockDB) GetGamesPlayed() ([]*models.GamesPlayed, error) {
	return []*models.GamesPlayed{
		{
			ID:          1,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			DeletedAt:   nil, // Indicates the game is not deleted
			Title:       "Epic Adventure",
			Genre:       models.GameGenreUndefined,
			Rating:      4,
			Description: "Embark on an epic journey through uncharted lands.",
		},
		{
			ID:          2,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			DeletedAt:   nil,
			Title:       "Mystic Quest",
			Genre:       models.GameGenreTableTop,
			Rating:      4,
			Description: "Dive into the mystic world and unravel its secrets.",
		},
		{
			ID:          3,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			DeletedAt:   nil,
			Title:       "Mystic Quest",
			Genre:       models.GameGenreStrategy,
			Rating:      4,
			Description: "Dive into the mystic world and unravel its secrets.",
		},
		{
			ID:          4,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			DeletedAt:   nil,
			Title:       "Mystic Quest",
			Genre:       models.GameGenreRpg,
			Rating:      4,
			Description: "Dive into the mystic world and unravel its secrets.",
		},
		{
			ID:          5,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			DeletedAt:   nil,
			Title:       "Mystic Quest",
			Genre:       models.GameGenreTableTop,
			Rating:      4,
			Description: "Dive into the mystic world and unravel its secrets.",
		},
		{
			ID:          6,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			DeletedAt:   nil,
			Title:       "Mystic Quest",
			Genre:       models.GameGenreTableTop,
			Rating:      4,
			Description: "Dive into the mystic world and unravel its secrets.",
		},
	}, nil
}
