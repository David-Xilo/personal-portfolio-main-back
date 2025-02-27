package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"safehouse-main-back/src/internal/models"

	"safehouse-main-back/src/internal/service"
)

type TechController struct {
	db *gorm.DB
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
	c.JSON(http.StatusOK, gin.H{"message": "This is the Tech Intro screen."})
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
	projects := tc.getProjects()
	c.JSON(http.StatusOK, gin.H{"message": projects})
}

func (tc *TechController) getProjectsRequest(w http.ResponseWriter) {
	projects := tc.getProjects()
	service.GetJSONData(w, projects)
}

func (tc *TechController) getProjects() []*models.TechProjectsDTO {
	var projects []*models.TechProjects

	if err := tc.db.Limit(10).Find(&projects).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []*models.TechProjectsDTO{}
		}
		panic(err)
	}

	projectsDTOList := models.ToTechProjectsDTOList(projects)

	return projectsDTOList
}
