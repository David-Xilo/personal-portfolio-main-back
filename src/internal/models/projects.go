package models

type RepositoriesDTO struct {
	Title       string     `json:"title"`
	Genre       GameGenres `json:"genre"`
	Rating      int        `json:"rating"`
	Description string     `json:"description"`
	LinkToGit   string     `json:"link_to_git"`
	LinkToStore string     `json:"link_to_store"`
}

func GameProjectsToProjectsDTO(gameProject *GameRepositories) *RepositoriesDTO {
	return &RepositoriesDTO{
		Title:       gameProject.Title,
		Genre:       gameProject.Genre,
		Rating:      gameProject.Rating,
		Description: gameProject.Description,
		LinkToGit:   gameProject.LinkToGit,
		LinkToStore: gameProject.LinkToStore,
	}
}

func TechProjectsToProjectsDTO(techProject *TechRepositories) *RepositoriesDTO {
	return &RepositoriesDTO{
		Title:       techProject.Title,
		Description: techProject.Description,
		LinkToGit:   techProject.LinkToGit,
	}
}

func FinanceProjectsToProjectsDTO(financeProjects *FinanceRepositories) *RepositoriesDTO {
	return &RepositoriesDTO{
		Title:       financeProjects.Title,
		Description: financeProjects.Description,
		LinkToGit:   financeProjects.LinkToGit,
	}
}
