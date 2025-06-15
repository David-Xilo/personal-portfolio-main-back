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
	router.GET("/games/intro", gc.handleIntro)
	router.GET("/games/projects", gc.handleProjects)
}

// @Summary Get introduction to the games section
// @Description Returns a brief introduction to the games section
// @Tags games
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]string
// @Router /games/intro [get]
func (gc *GamesController) handleIntro(c *gin.Context) {
	gamesIntroMessage := "This section is all about games.\n " +
		"Since I was a kid I've loved video games, in fact, they were one of the main reasons I got interested in computers.\n " +
		"This page might be empty (it is for sure emptier than I'd like), but I'll keep adding new games I've made here. " +
		"Who knows — maybe one day you'll recognize one of the names!"
	c.JSON(http.StatusOK, gin.H{"message": gamesIntroMessage})
}

// @Summary Get projects related to games
// @Description Returns a list of projects related to games
// @Tags games
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.GamesDTO
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
