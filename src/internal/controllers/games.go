package controllers

import (
	"errors"
	"gorm.io/gorm"
	"net/http"
	"safehouse-main-back/src/internal/models"
	"safehouse-main-back/src/internal/service"
)

type GamesController struct {
	db *gorm.DB
}

// @Summary Games list
// @Description Returns a list of games
// @Tags games
// @Produce json
// @Success 200 {object} []string
// @Router /games [get]
func (gc *GamesController) HandleGamesRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		gc.handleGetRequest(w, r)
	case http.MethodPost:
		gc.handlePostRequest(w, r)
	default:
		// Return an error for unsupported HTTP methods
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (gc *GamesController) handleGetRequest(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/intro":
		service.GetJSONSimpleStringMessage(w, "Welcome to the games intro!")
	case "/news":
		service.GetNewsByGenre(w, models.NewsGenreGaming, gc.db)
	case "/news/topic-of-the-season":
		service.GetTopicOfTheSeasonByGenre(w, models.NewsGenreGaming, gc.db)
	case "/genres":
		getGameGenres(w)
	case "/projects":
		gc.getGamesRequest(w)
	default:
		// Handle other /games routes or return 404 error
		http.NotFound(w, r)
	}
}

func getGameGenres(w http.ResponseWriter) {
	genres := models.GetAllGameGenres()
	service.GetJSONData(w, genres)
}

func (gc *GamesController) getGamesRequest(w http.ResponseWriter) {
	games := getGames(gc.db)
	service.GetJSONData(w, games)
}

func getGames(db *gorm.DB) []*models.Games {
	var games []*models.Games

	if err := db.
		Order("created_at desc").
		Limit(5).
		Find(&games).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []*models.Games{}
		}
		panic(err)
	}

	return games
}

func (gc *GamesController) handlePostRequest(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/games/filter":
		gc.getGamesFilteredByGenre(w, r)
	default:
		http.NotFound(w, r)
	}
}

func (gc *GamesController) getGamesFilteredByGenre(w http.ResponseWriter, r *http.Request) {
	result, err := gc.filterGames(w, r)
	if err != nil {
		panic(err)
	}
	service.GetJSONData(w, result)
}

func (gc *GamesController) filterGames(w http.ResponseWriter, r *http.Request) ([]*models.Games, error) {

	var results []*models.Games

	return results, nil
}
