package models

import (
	"time"
)

type TechProjects struct {
	ID          uint       `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	LinkToGit   string     `json:"link_to_git"`
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
