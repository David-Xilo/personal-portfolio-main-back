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

	TechProjects    []TechRepositories    `json:"tech_projects,omitempty" gorm:"foreignKey:ProjectGroupID"`
	GameProjects    []GameRepositories    `json:"game_projects,omitempty" gorm:"foreignKey:ProjectGroupID"`
	FinanceProjects []FinanceRepositories `json:"finance_projects,omitempty" gorm:"foreignKey:ProjectGroupID"`
}

type ProjectGroupsDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ProjectType string `json:"project_type"`

	Repositories []*RepositoriesDTO `json:"repositories,omitempty"`
}

func ToProjectGroupsDTO(projectGroup *ProjectGroups) *ProjectGroupsDTO {
	var repositoriesDTOList []*RepositoriesDTO
	for _, techProject := range projectGroup.TechProjects {
		dto := TechProjectsToProjectsDTO(&techProject)
		repositoriesDTOList = append(repositoriesDTOList, dto)
	}

	for _, gameProject := range projectGroup.GameProjects {
		dto := GameProjectsToProjectsDTO(&gameProject)
		repositoriesDTOList = append(repositoriesDTOList, dto)
	}

	for _, financeProject := range projectGroup.FinanceProjects {
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
	var projectGroupsDTOList []*ProjectGroupsDTO
	for _, techProject := range projectGroups {
		dto := ToProjectGroupsDTO(techProject)
		projectGroupsDTOList = append(projectGroupsDTOList, dto)
	}
	return projectGroupsDTOList
}
