package controllers

import (
	"gorm.io/gorm"
	"net/http"
	"safehouse-main-back/src/internal/models"
	"safehouse-main-back/src/internal/service"
)

type FinanceController struct {
	db *gorm.DB
}

func (fc *FinanceController) HandleFinanceRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		fc.handleGetRequest(w, r)
	default:
		// Return an error for unsupported HTTP methods
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (fc *FinanceController) handleGetRequest(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/intro":
		service.GetJSONSimpleStringMessage(w, "This is the finance Intro screen.")
	case "/news":
		service.GetNewsByGenre(w, models.NewsGenreFinance, fc.db)
	case "/news/topic-of-the-season":
		service.GetTopicOfTheSeasonByGenre(w, models.NewsGenreFinance, fc.db)
	case "/timeframes":
		service.GetJSONSimpleStringMessage(w, "This is the finance studies screen.")
	case "/assets":
		service.GetJSONSimpleStringMessage(w, "This is the finance studies screen.")
	case "/data/{asset}/{timeframe}":
		service.GetJSONSimpleStringMessage(w, "This is the finance studies screen.")
	default:
		// Handle other /games routes or return 404 error
		http.NotFound(w, r)
	}
}
