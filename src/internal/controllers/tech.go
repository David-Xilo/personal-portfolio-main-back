package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"safehouse-main-back/src/internal/database"
	"safehouse-main-back/src/internal/models"

	"safehouse-main-back/src/internal/service"
)

type TechController struct {
	db database.Database
}

func (tc *TechController) RegisterRoutes(router *gin.Engine) {
	router.GET("/tech/intro", tc.handleIntro)
	router.GET("/tech/projects", tc.handleProjects)
}

// @Summary Get introduction to the tech section
// @Description Returns an introductory message for the tech section
// @Tags tech
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]string
// @Router /tech/intro [get]
func (tc *TechController) handleIntro(c *gin.Context) {
	techIntroMessage := "Technology is always improving — and so am I.\n " +
		"Technology is my bread and butter. I've grown around it, worked with it, and played with it — " +
		"It's part of my personality and I try to learn and explore it as much as I can.\n " +
		"I'll post my personal projects here as a way to keep myself accountable.\n" +
		" Have fun exploring them!"
	c.JSON(http.StatusOK, gin.H{"message": techIntroMessage})
}

// @Summary Get projects related to tech
// @Description Returns a list of tech-related projects
// @Tags tech
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.TechProjectsDTO
// @Failure 404 {object} map[string]string
// @Router /tech/projects [get]
func (tc *TechController) handleProjects(c *gin.Context) {
	projects, _ := tc.db.GetTechProjects()
	c.JSON(http.StatusOK, gin.H{"message": projects})
}

func (tc *TechController) getProjectsRequest(w http.ResponseWriter) {
	projects, _ := tc.db.GetTechProjects()
	projectsDTOList := models.ToTechProjectsDTOList(projects)
	service.GetJSONData(w, projectsDTOList)
}
