package endpoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"safehouse-main-back/src/internal/database"
	"safehouse-main-back/src/internal/models"
)

type FinanceController struct {
	db database.Database
}

func NewFinanceController(db database.Database) *FinanceController {
	return &FinanceController{
		db: db,
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
func (fc *FinanceController) handleProjects(c *gin.Context) {
	projects, _ := fc.db.GetProjects(models.ProjectTypeFinance)
	projectsDTOList := models.ToProjectGroupsDTOList(projects)
	c.JSON(http.StatusOK, gin.H{"message": projectsDTOList})
}
