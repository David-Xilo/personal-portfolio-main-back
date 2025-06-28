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

type FinanceController struct {
	db     database.Database
	config configuration.Config
}

func NewFinanceController(db database.Database, config configuration.Config) *FinanceController {
	return &FinanceController{
		db:     db,
		config: config,
	}
}

func (fc *FinanceController) RegisterRoutes(router *gin.Engine) {
	router.GET("/finance/projects", fc.handleProjects)
}

// @Summary Get projects related to finance
// @Description Returns a list of finance-related projects
// @Tags finance
// @Accept  json
// @Produce  json
// @Success 200 {array} []models.ProjectGroupsDTO
// @Failure 404 {object} map[string]string
// @Router /finance/projects [get]
func (fc *FinanceController) handleProjects(ctx *gin.Context) {
	projects, err := middleware.WithTimeout(ctx.Request.Context(), fc.config.DatabaseTimeout, func(dbCtx context.Context) ([]*models.ProjectGroups, error) {
		return fc.db.GetProjects(models.ProjectTypeGame)
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
	projectsDTOList := models.ToProjectGroupsDTOList(projects)
	ctx.JSON(http.StatusOK, gin.H{"message": projectsDTOList})
}
