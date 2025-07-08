package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTechRepositoriesStruct(t *testing.T) {
	now := time.Date(2023, time.January, 1, 12, 0, 0, 0, time.UTC)
	deletedAt := time.Date(2023, time.January, 1, 12, 0, 0, 0, time.UTC).Add(1 * time.Hour)

	techRepo := TechRepositories{
		ID:             123,
		CreatedAt:      now,
		UpdatedAt:      now,
		DeletedAt:      &deletedAt,
		ProjectGroupID: 456,
		Title:          "Tech Repository",
		Description:    "A technical repository",
		LinkToGit:      "https://github.com/user/tech-repo",
	}

	assert.Equal(t, uint(123), techRepo.ID)
	assert.Equal(t, uint(456), techRepo.ProjectGroupID)
	assert.Equal(t, "Tech Repository", techRepo.Title)
	assert.Equal(t, "A technical repository", techRepo.Description)
	assert.Equal(t, "https://github.com/user/tech-repo", techRepo.LinkToGit)
	assert.NotNil(t, techRepo.DeletedAt)
	assert.Equal(t, deletedAt, *techRepo.DeletedAt)
}

func TestTechRepositoriesDTOStruct(t *testing.T) {
	dto := TechRepositoriesDTO{
		Title:       "Tech DTO Repository",
		Description: "A technical DTO repository",
		LinkToGit:   "https://github.com/user/tech-dto-repo",
	}

	assert.Equal(t, "Tech DTO Repository", dto.Title)
	assert.Equal(t, "A technical DTO repository", dto.Description)
	assert.Equal(t, "https://github.com/user/tech-dto-repo", dto.LinkToGit)
}

func TestTechRepositories_ForeignKey(t *testing.T) {
	// Test that the foreign key relationship can be set
	techRepo := TechRepositories{
		ProjectGroupID: 789,
		Title:          "Test Tech Repo",
		Description:    "Test description",
		LinkToGit:      "https://github.com/test/repo",
	}

	assert.Equal(t, uint(789), techRepo.ProjectGroupID)
	assert.Equal(t, "Test Tech Repo", techRepo.Title)
}

func TestTechRepositories_EmptyFields(t *testing.T) {
	// Test with empty fields
	techRepo := TechRepositories{
		ID:             1,
		ProjectGroupID: 2,
		Title:          "",
		Description:    "",
		LinkToGit:      "",
	}

	assert.Equal(t, uint(1), techRepo.ID)
	assert.Equal(t, uint(2), techRepo.ProjectGroupID)
	assert.Equal(t, "", techRepo.Title)
	assert.Equal(t, "", techRepo.Description)
	assert.Equal(t, "", techRepo.LinkToGit)
}
