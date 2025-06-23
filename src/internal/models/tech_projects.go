package models

import (
	"time"
)

type TechRepositories struct {
	ID             uint       `json:"id"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at,omitempty"`
	ProjectGroupID uint       `json:"project_group_id"`
	Title          string     `json:"title"`
	Description    string     `json:"description"`
	LinkToGit      string     `json:"link_to_git"`

	//ProjectGroup ProjectGroups `json:"project_group,omitempty" gorm:"foreignKey:ProjectGroupID"`
}

type TechRepositoriesDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	LinkToGit   string `json:"link_to_git"`
}

//func ToTechProjectsDTO(techProject *TechRepositories) *TechRepositoriesDTO {
//	return &TechRepositoriesDTO{
//		Title:       techProject.Title,
//		Description: techProject.Description,
//		LinkToGit:   techProject.LinkToGit,
//	}
//}

//func ToTechProjectsDTOList(techProjects []*TechRepositories) []*TechRepositoriesDTO {
//	var techProjectsDTOList []*TechRepositoriesDTO
//	for _, techProject := range techProjects {
//		dto := ToTechProjectsDTO(techProject)
//		techProjectsDTOList = append(techProjectsDTOList, dto)
//	}
//	return techProjectsDTOList
//}
