package models

import (
	"time"
)

type ProjectGroups struct {
	ID          uint       `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	ProjectType string     `json:"project_type"`

	TechRepositories    []TechRepositories    `json:"tech_repositories,omitempty" gorm:"foreignKey:ProjectGroupID"`
	GameRepositories    []GameRepositories    `json:"game_repositories,omitempty" gorm:"foreignKey:ProjectGroupID"`
	FinanceRepositories []FinanceRepositories `json:"finance_repositories,omitempty" gorm:"foreignKey:ProjectGroupID"`
}

type ProjectGroupsDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ProjectType string `json:"project_type"`

	Repositories []*RepositoriesDTO `json:"repositories,omitempty"`
}

func ToProjectGroupsDTO(projectGroup *ProjectGroups) *ProjectGroupsDTO {
	repositoriesDTOList := make([]*RepositoriesDTO, 0)
	for _, techProject := range projectGroup.TechRepositories {
		dto := TechProjectsToProjectsDTO(&techProject)
		repositoriesDTOList = append(repositoriesDTOList, dto)
	}

	for _, gameProject := range projectGroup.GameRepositories {
		dto := GameProjectsToProjectsDTO(&gameProject)
		repositoriesDTOList = append(repositoriesDTOList, dto)
	}

	for _, financeProject := range projectGroup.FinanceRepositories {
		dto := FinanceProjectsToProjectsDTO(&financeProject)
		repositoriesDTOList = append(repositoriesDTOList, dto)
	}

	return &ProjectGroupsDTO{
		Title:        projectGroup.Title,
		Description:  projectGroup.Description,
		ProjectType:  projectGroup.ProjectType,
		Repositories: repositoriesDTOList,
	}
}

func ToProjectGroupsDTOList(projectGroups []*ProjectGroups) []*ProjectGroupsDTO {
	projectGroupsDTOList := make([]*ProjectGroupsDTO, 0)
	for _, techProject := range projectGroups {
		dto := ToProjectGroupsDTO(techProject)
		projectGroupsDTOList = append(projectGroupsDTOList, dto)
	}
	return projectGroupsDTOList
}
