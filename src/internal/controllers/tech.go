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
	router.GET("/tech/news", tc.handleNews)
	router.GET("/tech/news/topic-of-the-season", tc.handleTopicOfTheSeason)
	router.GET("/tech/projects", tc.handleProjects)
}

func (tc *TechController) handleIntro(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "This is the Tech Intro screen."})
}

func (tc *TechController) handleNews(c *gin.Context) {
	service.GetNewsByGenre(c, models.NewsGenreTech, tc.db)
}

func (tc *TechController) handleTopicOfTheSeason(c *gin.Context) {
	service.GetTopicOfTheSeasonByGenre(c, models.NewsGenreTech, tc.db)
}

func (tc *TechController) handleProjects(c *gin.Context) {
	projects := tc.getProjects()
	c.JSON(http.StatusOK, gin.H{"message": projects})
}

func (tc *TechController) getProjectsRequest(w http.ResponseWriter) {
	projects := tc.getProjects()
	service.GetJSONData(w, projects)
}

func (tc *TechController) getProjects() []*models.TechProjects {
	var projects []*models.TechProjects

	if err := tc.db.Limit(10).Find(&projects).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []*models.TechProjects{}
		}
		panic(err)
	}
	return projects
}
