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

func (gc *GamesController) handleIntro(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "This is the Tech Intro screen."})
}

func (gc *GamesController) handleNews(c *gin.Context) {
	service.GetNewsByGenre(c, models.NewsGenreTech, gc.db)
}

func (gc *GamesController) handleTopicOfTheSeason(c *gin.Context) {
	service.GetTopicOfTheSeasonByGenre(c, models.NewsGenreTech, gc.db)
}

func (gc *GamesController) handleProjects(c *gin.Context) {
	games := getGames(gc.db)
	c.JSON(http.StatusOK, gin.H{"message": games})
}

func (gc *GamesController) handleGenres(c *gin.Context) {
	genres := models.GetAllGameGenres()
	c.JSON(http.StatusOK, gin.H{"message": genres})
}

func (gc *GamesController) handleGamesFiltered(c *gin.Context) {
	var filter GamesFilter

	if err := c.ShouldBind(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid filter"})
		return
	}

	games := getGamesFiltered(gc.db, filter)
	c.JSON(http.StatusOK, gin.H{"message": games})
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

func getGamesFiltered(db *gorm.DB, filter GamesFilter) []*models.Games {
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
			return []*models.Games{}
		}
		panic(err)
	}

	return games
}
