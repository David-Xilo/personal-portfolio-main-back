package models

import (
	"time"
)

type FinanceProjects struct {
	ID             uint       `json:"id"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at,omitempty"`
	ProjectGroupID uint       `json:"project_group_id"`
	Title          string     `json:"title"`
	Description    string     `json:"description"`
	LinkToGit      string     `json:"link_to_git"`

	ProjectGroup ProjectGroups `json:"project_group,omitempty" gorm:"foreignKey:ProjectGroupID"`
}

type FinanceProjectsDTO struct {
	//ProjectGroupID uint       `json:"project_group_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	LinkToGit   string `json:"link_to_git"`

	//ProjectGroup ProjectGroups `json:"project_group,omitempty"`
}

func ToFinanceProjectsDTO(financeProject *FinanceProjects) *FinanceProjectsDTO {
	return &FinanceProjectsDTO{
		Title:       financeProject.Title,
		Description: financeProject.Description,
		LinkToGit:   financeProject.LinkToGit,
		//ProjectGroup: financeProject.ProjectGroup,
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
