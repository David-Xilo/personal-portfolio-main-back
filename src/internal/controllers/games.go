package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"safehouse-main-back/src/internal/models"
	"safehouse-main-back/src/internal/service"
)

type GamesController struct {
	db *gorm.DB
}

type GamesFilter struct {
	Genre  string `json:"genre"`
	Rating *int   `json:"rating"`
}

func (gc *GamesController) RegisterRoutes(router *gin.Engine) {
	router.GET("/games/intro", gc.handleIntro)
	router.GET("/games/news", gc.handleNews)
	router.GET("/games/news/topic-of-the-season", gc.handleTopicOfTheSeason)
	router.GET("/games/genres", gc.handleGenres)
	router.GET("/games/projects", gc.handleProjects)

	router.POST("/games/filter", gc.handleGamesFiltered)
}

// @Summary Get introduction to the games section
// @Description Returns a brief introduction to the games section
// @Tags games
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]string
// @Router /games/intro [get]
func (gc *GamesController) handleIntro(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "This is the Tech Intro screen."})
}

// @Summary Get news related to games
// @Description Returns a list of news related to games
// @Tags games
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.News
// @Failure 404 {object} map[string]string
// @Router /games/news [get]
func (gc *GamesController) handleNews(c *gin.Context) {
	service.GetNewsByGenre(c, models.NewsGenreTech, gc.db)
}

// @Summary Get topic of the season for games
// @Description Returns the topic of the season for the games genre
// @Tags games
// @Accept  json
// @Produce  json
// @Success 200 {object} service.NewsWithTopic
// @Failure 404 {object} map[string]string
// @Router /games/news/topic-of-the-season [get]
func (gc *GamesController) handleTopicOfTheSeason(c *gin.Context) {
	service.GetTopicOfTheSeasonByGenre(c, models.NewsGenreTech, gc.db)
}

// @Summary Get projects related to games
// @Description Returns a list of projects related to games
// @Tags games
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.Games
// @Failure 404 {object} map[string]string
// @Router /games/projects [get]
func (gc *GamesController) handleProjects(c *gin.Context) {
	games := getGames(gc.db)
	c.JSON(http.StatusOK, gin.H{"message": games})
}

// @Summary Get a list of game genres
// @Description Returns a list of all available game genres
// @Tags games
// @Accept  json
// @Produce  json
// @Success 200 {object} []string
// @Failure 404 {object} map[string]string
// @Router /games/genres [get]
func (gc *GamesController) handleGenres(c *gin.Context) {
	genres := models.GetAllGameGenres()
	c.JSON(http.StatusOK, gin.H{"message": genres})
}

// @Summary Filter games by genre and rating
// @Description Filters the games based on genre and rating parameters
// @Tags games
// @Accept  json
// @Produce  json
// @Param filter body GamesFilter true "Filter parameters for games"
// @Success 200 {object} []models.Games
// @Failure 400 {object} map[string]string
// @Router /games/filter [post]
func (gc *GamesController) handleGamesFiltered(c *gin.Context) {
	var filter GamesFilter

	if err := c.ShouldBind(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid filter"})
		return
	}

	games := getGamesFiltered(gc.db, filter)
	c.JSON(http.StatusOK, gin.H{"message": games})
}

func getGames(db *gorm.DB) []*models.GamesDTO {
	var games []*models.Games

	if err := db.
		Order("created_at desc").
		Limit(5).
		Find(&games).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []*models.GamesDTO{}
		}
		panic(err)
	}

	gamesDTOList := models.ToGamesListDTO(games)

	return gamesDTOList
}

func getGamesFiltered(db *gorm.DB, filter GamesFilter) []*models.GamesDTO {
	var games []*models.Games

	genre := filter.Genre
	rating := filter.Rating

	query := db.Order("created_at desc").Limit(5)

	if genre != "" {
		query = query.Where("genre = ?", genre)
	}

	if rating != nil && *rating >= 0 {
		query = query.Where("rating = ?", *rating)
	}

	if err := query.Find(&games).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []*models.GamesDTO{}
		}
		panic(err)
	}

	gamesDTOList := models.ToGamesListDTO(games)

	return gamesDTOList
}
