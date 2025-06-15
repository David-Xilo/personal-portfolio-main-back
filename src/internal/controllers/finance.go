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
	financeIntroMessage := "Finance is one of my passions. I read about it, study it, and pay attention to it in my daily life.\n " +
		"Since I was young, I’ve followed the stock market and economic news. " +
		"During my master’s thesis, I finally combined my passions for technology and finance—and I loved it.\n " +
		"Now, I’ll post my finance-related personal projects here, along with any certifications or extra courses I complete in my free time."
	c.JSON(http.StatusOK, gin.H{"message": financeIntroMessage})
}

// @Summary Get projects related to finance
// @Description Returns a list of finance-related projects
// @Tags finance
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.FinanceProjectsDTO
// @Failure 404 {object} map[string]string
// @Router /finance/projects [get]
func (fc *FinanceController) handleProjects(c *gin.Context) {
	projects, _ := fc.db.GetFinanceProjects()
	c.JSON(http.StatusOK, gin.H{"message": projects})
}

func (fc *FinanceController) getProjectsRequest(w http.ResponseWriter) {
	projects, _ := fc.db.GetFinanceProjects()
	projectsDTOList := models.ToFinanceProjectsDTOList(projects)
	service.GetJSONData(w, projectsDTOList)
}
