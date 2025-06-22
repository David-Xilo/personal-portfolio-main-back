package models

type ProjectsDTO struct {
	Title       string     `json:"title"`
	Genre       GameGenres `json:"genre"`
	Rating      int        `json:"rating"`
	Description string     `json:"description"`
	LinkToGit   string     `json:"link_to_git"`
	LinkToStore string     `json:"link_to_store"`
}

func GameProjectsToProjectsDTO(gameProject *GameProjects) *ProjectsDTO {
	return &ProjectsDTO{
		Title:       gameProject.Title,
		Genre:       gameProject.Genre,
		Rating:      gameProject.Rating,
		Description: gameProject.Description,
		LinkToGit:   gameProject.LinkToGit,
		LinkToStore: gameProject.LinkToStore,
	}
}

func TechProjectsToProjectsDTO(techProject *TechProjects) *ProjectsDTO {
	return &ProjectsDTO{
		Title:       techProject.Title,
		Description: techProject.Description,
		LinkToGit:   techProject.LinkToGit,
	}
}

func FinanceProjectsToProjectsDTO(financeProjects *FinanceProjects) *ProjectsDTO {
	return &ProjectsDTO{
		Title:       financeProjects.Title,
		Description: financeProjects.Description,
		LinkToGit:   financeProjects.LinkToGit,
	}
}

func GameProjectsToProjectsDTOList(gameProjects []*GameProjects) []*ProjectsDTO {
	var projectsDTO []*ProjectsDTO
	for _, gameProject := range gameProjects {
		dto := GameProjectsToProjectsDTO(gameProject)
		projectsDTO = append(projectsDTO, dto)
	}
	return projectsDTO
}

func TechProjectsToProjectsDTOList(techProjects []*TechProjects) []*ProjectsDTO {
	var projectsDTO []*ProjectsDTO
	for _, techProject := range techProjects {
		dto := TechProjectsToProjectsDTO(techProject)
		projectsDTO = append(projectsDTO, dto)
	}
	return projectsDTO
}

func FinanceProjectsToProjectsDTOList(financeProjects []*FinanceProjects) []*ProjectsDTO {
	var projectsDTO []*ProjectsDTO
	for _, financeProject := range financeProjects {
		dto := FinanceProjectsToProjectsDTO(financeProject)
		projectsDTO = append(projectsDTO, dto)
	}
	return projectsDTO
}
