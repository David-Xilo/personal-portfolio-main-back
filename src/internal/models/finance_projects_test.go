package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFinanceRepositoriesStruct(t *testing.T) {
	now := time.Now()
	deletedAt := time.Now().Add(1 * time.Hour)

	financeRepo := FinanceRepositories{
		ID:             123,
		CreatedAt:      now,
		UpdatedAt:      now,
		DeletedAt:      &deletedAt,
		ProjectGroupID: 456,
		Title:          "Finance Repository",
		Description:    "A financial repository",
		LinkToGit:      "https://github.com/user/finance-repo",
	}

	assert.Equal(t, uint(123), financeRepo.ID)
	assert.Equal(t, uint(456), financeRepo.ProjectGroupID)
	assert.Equal(t, "Finance Repository", financeRepo.Title)
	assert.Equal(t, "A financial repository", financeRepo.Description)
	assert.Equal(t, "https://github.com/user/finance-repo", financeRepo.LinkToGit)
	assert.NotNil(t, financeRepo.DeletedAt)
	assert.Equal(t, deletedAt, *financeRepo.DeletedAt)
}

func TestFinanceRepositoriesDTOStruct(t *testing.T) {
	dto := FinanceRepositoriesDTO{
		Title:       "Finance DTO Repository",
		Description: "A financial DTO repository",
		LinkToGit:   "https://github.com/user/finance-dto-repo",
	}

	assert.Equal(t, "Finance DTO Repository", dto.Title)
	assert.Equal(t, "A financial DTO repository", dto.Description)
	assert.Equal(t, "https://github.com/user/finance-dto-repo", dto.LinkToGit)
}

func TestFinanceRepositories_ForeignKey(t *testing.T) {
	// Test that the foreign key relationship can be set
	financeRepo := FinanceRepositories{
		ProjectGroupID: 789,
		Title:          "Test Finance Repo",
		Description:    "Test finance description",
		LinkToGit:      "https://github.com/test/finance",
	}

	assert.Equal(t, uint(789), financeRepo.ProjectGroupID)
	assert.Equal(t, "Test Finance Repo", financeRepo.Title)
	assert.Equal(t, "Test finance description", financeRepo.Description)
	assert.Equal(t, "https://github.com/test/finance", financeRepo.LinkToGit)
}

func TestFinanceRepositories_EmptyFields(t *testing.T) {
	// Test with empty fields
	financeRepo := FinanceRepositories{
		ID:             1,
		ProjectGroupID: 2,
		Title:          "",
		Description:    "",
		LinkToGit:      "",
	}

	assert.Equal(t, uint(1), financeRepo.ID)
	assert.Equal(t, uint(2), financeRepo.ProjectGroupID)
	assert.Equal(t, "", financeRepo.Title)
	assert.Equal(t, "", financeRepo.Description)
	assert.Equal(t, "", financeRepo.LinkToGit)
}