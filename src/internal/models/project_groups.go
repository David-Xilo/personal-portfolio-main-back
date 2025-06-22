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

	TechProjects    []TechProjects    `json:"tech_projects,omitempty" gorm:"foreignKey:ProjectGroupID"`
	GameProjects    []GameProjects    `json:"game_projects,omitempty" gorm:"foreignKey:ProjectGroupID"`
	FinanceProjects []FinanceProjects `json:"finance_projects,omitempty" gorm:"foreignKey:ProjectGroupID"`
}

type ProjectGroupsDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ProjectType string `json:"project_type"`

	Projects []*ProjectsDTO `json:"projects,omitempty"`
}

func ToProjectGroupsDTO(projectGroup *ProjectGroups) *ProjectGroupsDTO {
	var projectsDTOList []*ProjectsDTO
	for _, techProject := range projectGroup.TechProjects {
		dto := TechProjectsToProjectsDTO(&techProject)
		projectsDTOList = append(projectsDTOList, dto)
	}

	for _, gameProject := range projectGroup.GameProjects {
		dto := GameProjectsToProjectsDTO(&gameProject)
		projectsDTOList = append(projectsDTOList, dto)
	}

	for _, financeProject := range projectGroup.FinanceProjects {
		dto := FinanceProjectsToProjectsDTO(&financeProject)
		projectsDTOList = append(projectsDTOList, dto)
	}

	return &ProjectGroupsDTO{
		Title:       projectGroup.Title,
		Description: projectGroup.Description,
		ProjectType: projectGroup.ProjectType,
		Projects:    projectsDTOList,
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
