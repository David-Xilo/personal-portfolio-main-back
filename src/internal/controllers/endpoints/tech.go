package endpoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"safehouse-main-back/src/internal/database"
	"safehouse-main-back/src/internal/models"
)

type TechController struct {
	db database.Database
}

func NewTechController(db database.Database) *TechController {
	return &TechController{
		db: db,
	}
}

func (tc *TechController) RegisterRoutes(router *gin.Engine) {
	router.GET("/tech/projects", tc.handleProjects)
}

// @Summary Get projects related to tech
// @Description Returns a list of tech-related projects
// @Tags tech
// @Accept  json
// @Produce  json
// @Success 200 {array} []models.ProjectGroupsDTO
// @Failure 404 {object} map[string]string
// @Router /tech/projects [get]
func (tc *TechController) handleProjects(c *gin.Context) {
	projects, _ := tc.db.GetProjects(models.ProjectTypeTech)
	projectsDTOList := models.ToProjectGroupsDTOList(projects)
	c.JSON(http.StatusOK, gin.H{"message": projectsDTOList})
}
