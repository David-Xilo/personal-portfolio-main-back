package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGameRepositoriesStruct(t *testing.T) {
	now := time.Date(2023, time.January, 1, 12, 0, 0, 0, time.UTC)
	deletedAt := time.Date(2023, time.January, 1, 12, 0, 0, 0, time.UTC).Add(1 * time.Hour)

	gameRepo := GameRepositories{
		ID:             123,
		CreatedAt:      now,
		UpdatedAt:      now,
		DeletedAt:      &deletedAt,
		ProjectGroupID: 456,
		Title:          "Game Repository",
		Genre:          "Action", // Assuming GameGenres is a string type
		Rating:         5,
		Description:    "An action game repository",
		LinkToGit:      "https://github.com/user/game-repo",
		LinkToStore:    "https://store.com/game",
	}

	assert.Equal(t, uint(123), gameRepo.ID)
	assert.Equal(t, uint(456), gameRepo.ProjectGroupID)
	assert.Equal(t, "Game Repository", gameRepo.Title)
	assert.Equal(t, GameGenres("Action"), gameRepo.Genre)
	assert.Equal(t, 5, gameRepo.Rating)
	assert.Equal(t, "An action game repository", gameRepo.Description)
	assert.Equal(t, "https://github.com/user/game-repo", gameRepo.LinkToGit)
	assert.Equal(t, "https://store.com/game", gameRepo.LinkToStore)
	assert.NotNil(t, gameRepo.DeletedAt)
	assert.Equal(t, deletedAt, *gameRepo.DeletedAt)
}

func TestGameRepositories_ForeignKey(t *testing.T) {
	// Test that the foreign key relationship can be set
	gameRepo := GameRepositories{
		ProjectGroupID: 789,
		Title:          "Test Game Repo",
		Genre:          "RPG",
		Rating:         4,
		Description:    "Test game description",
		LinkToGit:      "https://github.com/test/game",
		LinkToStore:    "https://store.com/test-game",
	}

	assert.Equal(t, uint(789), gameRepo.ProjectGroupID)
	assert.Equal(t, "Test Game Repo", gameRepo.Title)
	assert.Equal(t, GameGenres("RPG"), gameRepo.Genre)
	assert.Equal(t, 4, gameRepo.Rating)
}

func TestGameRepositories_EmptyFields(t *testing.T) {
	// Test with empty fields
	gameRepo := GameRepositories{
		ID:             1,
		ProjectGroupID: 2,
		Title:          "",
		Genre:          "",
		Rating:         0,
		Description:    "",
		LinkToGit:      "",
		LinkToStore:    "",
	}

	assert.Equal(t, uint(1), gameRepo.ID)
	assert.Equal(t, uint(2), gameRepo.ProjectGroupID)
	assert.Equal(t, "", gameRepo.Title)
	assert.Equal(t, GameGenres(""), gameRepo.Genre)
	assert.Equal(t, 0, gameRepo.Rating)
	assert.Equal(t, "", gameRepo.Description)
	assert.Equal(t, "", gameRepo.LinkToGit)
	assert.Equal(t, "", gameRepo.LinkToStore)
}

func TestGameRepositories_RatingBounds(t *testing.T) {
	// Test rating bounds
	gameRepo := GameRepositories{
		Title:  "Test Game",
		Rating: 10, // High rating
	}
	assert.Equal(t, 10, gameRepo.Rating)

	gameRepo.Rating = -1 // Negative rating
	assert.Equal(t, -1, gameRepo.Rating)

	gameRepo.Rating = 0 // Zero rating
	assert.Equal(t, 0, gameRepo.Rating)
}
