package models

import (
	"time"
)

type TechProjects struct {
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

type TechProjectsDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	LinkToGit   string `json:"link_to_git"`
}

func ToTechProjectsDTO(techProject *TechProjects) *TechProjectsDTO {
	return &TechProjectsDTO{
		Title:       techProject.Title,
		Description: techProject.Description,
		LinkToGit:   techProject.LinkToGit,
	}
}

func ToTechProjectsDTOList(techProjects []*TechProjects) []*TechProjectsDTO {
	var techProjectsDTOList []*TechProjectsDTO
	for _, techProject := range techProjects {
		dto := ToTechProjectsDTO(techProject)
		techProjectsDTOList = append(techProjectsDTOList, dto)
	}
	return techProjectsDTOList
}
