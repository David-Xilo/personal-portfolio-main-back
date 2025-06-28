package endpoints

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"net/http"
	configuration "safehouse-main-back/src/internal/config"
	"safehouse-main-back/src/internal/database"
	"safehouse-main-back/src/internal/middleware"
	"safehouse-main-back/src/internal/models"
)

type GamesController struct {
	db     database.Database
	config configuration.Config
}

func NewGamesController(db database.Database, config configuration.Config) *GamesController {
	return &GamesController{
		db:     db,
		config: config,
	}
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
// @Success 200 {array} []models.ProjectGroupsDTO
// @Failure 404 {object} map[string]string
// @Router /games/projects [get]
func (gc *GamesController) handleProjects(ctx *gin.Context) {
	games, err := middleware.WithTimeout(ctx.Request.Context(), gc.config.DatabaseTimeout, func(dbCtx context.Context) ([]*models.ProjectGroups, error) {
		return gc.db.GetProjects(models.ProjectTypeGame)
	})
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			ctx.JSON(http.StatusRequestTimeout, gin.H{"error": "Database timeout"})
			return
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNoContent, gin.H{"error": "No contact found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	projectsDTOList := models.ToProjectGroupsDTOList(games)
	ctx.JSON(http.StatusOK, gin.H{"message": projectsDTOList})
}

// @Summary Get projects related to games
// @Description Returns a list of projects related to games
// @Tags games
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.GamesPlayedDTO
// @Failure 404 {object} map[string]string
// @Router /games/projects [get]
func (gc *GamesController) handleGamesPlayedCarousel(ctx *gin.Context) {
	gamesPlayed, err := gc.db.GetGamesPlayed()
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			ctx.JSON(http.StatusRequestTimeout, gin.H{"error": "Database timeout"})
			return
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNoContent, gin.H{"error": "No contact found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	firstFive := gamesPlayed[:min(len(gamesPlayed), 5)]
	gamesPlayedDTOList := models.ToGamesPlayedListDTO(firstFive)
	ctx.JSON(http.StatusOK, gin.H{"message": gamesPlayedDTOList})
}
