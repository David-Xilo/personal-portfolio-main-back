package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"safehouse-main-back/src/internal/database"
	"safehouse-main-back/src/internal/models"
	"safehouse-main-back/src/internal/service"
)

type FinanceController struct {
	db database.Database
}

func (fc *FinanceController) RegisterRoutes(router *gin.Engine) {
	router.GET("/finance/intro", fc.handleIntro)
	router.GET("/finance/projects", fc.handleProjects)
}

// @Summary Get introduction to the finance section
// @Description Returns an introductory message for the finance section
// @Tags finance
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]string
// @Router /finance/intro [get]
func (fc *FinanceController) handleIntro(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "This is the Finance Intro screen."})
}

// @Summary Get projects related to tech
// @Description Returns a list of tech-related projects
// @Tags tech
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.TechProjectsDTO
// @Failure 404 {object} map[string]string
// @Router /tech/projects [get]
func (fc *FinanceController) handleProjects(c *gin.Context) {
	projects, _ := fc.db.GetFinanceProjects()
	c.JSON(http.StatusOK, gin.H{"message": projects})
}

func (fc *FinanceController) getProjectsRequest(w http.ResponseWriter) {
	projects, _ := fc.db.GetFinanceProjects()
	projectsDTOList := models.ToFinanceProjectsDTOList(projects)
	service.GetJSONData(w, projectsDTOList)
}
