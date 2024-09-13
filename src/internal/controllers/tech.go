package controllers

import (
	"errors"
	"gorm.io/gorm"
	"net/http"
	"safehouse-main-back/src/internal/models"

	"safehouse-main-back/src/internal/service"
)

type TechController struct {
	db *gorm.DB
}

func (tc *TechController) HandleTechRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		tc.handleGetRequest(w, r)
	default:
		// Return an error for unsupported HTTP methods
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (tc *TechController) handleGetRequest(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/intro":
		service.GetJSONSimpleStringMessage(w, "This is the Tech Intro screen.")
	case "/news":
		service.GetNewsByGenre(w, models.NewsGenreTech, tc.db)
	case "/news/topic-of-the-season":
		service.GetTopicOfTheSeasonByGenre(w, models.NewsGenreTech, tc.db)
	case "/projects":
		tc.getProjectsRequest(w)
	default:
		http.NotFound(w, r)
	}
}

func (tc *TechController) getProjectsRequest(w http.ResponseWriter) {
	projects := tc.getProjects()
	service.GetJSONData(w, projects)
}

func (tc *TechController) getProjects() []*models.TechProjects {
	var projects []*models.TechProjects

	if err := tc.db.Limit(10).Find(&projects).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []*models.TechProjects{}
		}
		panic(err)
	}
	return projects
}
