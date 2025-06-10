package models

import (
	"time"
)

type FinanceProjects struct {
	ID          uint       `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	LinkToGit   string     `json:"link_to_git"`
}

type FinanceProjectsDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	LinkToGit   string `json:"link_to_git"`
}

func ToFinanceProjectsDTO(techProject *FinanceProjects) *FinanceProjectsDTO {
	return &FinanceProjectsDTO{
		Title:       techProject.Title,
		Description: techProject.Description,
		LinkToGit:   techProject.LinkToGit,
	}
}

func ToFinanceProjectsDTOList(finProjects []*FinanceProjects) []*FinanceProjectsDTO {
	var finProjectsDTOList []*FinanceProjectsDTO
	for _, finProject := range finProjects {
		dto := ToFinanceProjectsDTO(finProject)
		finProjectsDTOList = append(finProjectsDTOList, dto)
	}
	return finProjectsDTOList
}
