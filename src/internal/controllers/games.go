package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"safehouse-main-back/src/internal/database"
	"safehouse-main-back/src/internal/models"
)

type GamesController struct {
	db database.Database
}

type GamesFilter struct {
	Genre string `json:"genre"`
}

func (gc *GamesController) RegisterRoutes(router *gin.Engine) {
	router.GET("/games/projects", gc.handleProjects)
	router.GET("/games/played/carousel", gc.handleGamesPlayedCarousel)
}

// @Summary Get projects related to games
// @Description Returns a list of projects related to games
// @Tags games
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.GameProjectsDTO
// @Failure 404 {object} map[string]string
// @Router /games/projects [get]
func (gc *GamesController) handleProjects(c *gin.Context) {
	games, _ := gc.db.GetGames()
	gamesDTOList := models.ToGamesListDTO(games)
	c.JSON(http.StatusOK, gin.H{"message": gamesDTOList})
}

// @Summary Get projects related to games
// @Description Returns a list of projects related to games
// @Tags games
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.GamesPlayedDTO
// @Failure 404 {object} map[string]string
// @Router /games/projects [get]
func (gc *GamesController) handleGamesPlayedCarousel(c *gin.Context) {
	gamesPlayed, _ := gc.db.GetGamesPlayed()
	firstFive := gamesPlayed[:min(len(gamesPlayed), 5)]
	gamesPlayedDTOList := models.ToGamesPlayedListDTO(firstFive)
	c.JSON(http.StatusOK, gin.H{"message": gamesPlayedDTOList})
}
