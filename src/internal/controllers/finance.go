package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"safehouse-main-back/src/internal/models"
	"safehouse-main-back/src/internal/service"
)

type FinanceController struct {
	db *gorm.DB
}

func (fc *FinanceController) RegisterRoutes(router *gin.Engine) {
	router.GET("/finance/intro", fc.handleIntro)
	router.GET("/finance/news", fc.handleNews)
	router.GET("/finance/news/topic-of-the-season", fc.handleTopicOfTheSeason)
	//router.GET("/finance/timeframes", fc.handleTimeframes)
	//router.GET("/finance/assets", fc.handleProjects)
	//router.GET("/finance/data/{asset}/{timeframe}", fc.handleData)
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

// @Summary Get finance news
// @Description Returns a list of news related to the finance genre
// @Tags finance
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.News
// @Failure 404 {object} map[string]string
// @Router /finance/news [get]
func (fc *FinanceController) handleNews(c *gin.Context) {
	service.GetNewsByGenre(c, models.NewsGenreTech, fc.db)
}

// @Summary Get topic of the season for finance
// @Description Returns the topic of the season for the finance genre
// @Tags finance
// @Accept  json
// @Produce  json
// @Success 200 {object} service.NewsWithTopic
// @Failure 404 {object} map[string]string
// @Router /finance/news/topic-of-the-season [get]
func (fc *FinanceController) handleTopicOfTheSeason(c *gin.Context) {
	service.GetTopicOfTheSeasonByGenre(c, models.NewsGenreTech, fc.db)
}

// Commented-out functions for timeframes, assets, and data can be annotated similarly once they are implemented
// @Summary Get available timeframes for finance
// @Description Returns a list of available timeframes for finance data
// @Tags finance
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.Timeframe
// @Router /finance/timeframes [get]
//func (fc *FinanceController) handleTimeframes(c *gin.Context) {
//	c.JSON(http.StatusOK, gin.H{"message": "This is Timeframes"})
//}

// @Summary Get list of finance assets
// @Description Returns a list of available finance assets
// @Tags finance
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.Asset
// @Router /finance/assets [get]
//func (fc *FinanceController) handleAssets(c *gin.Context) {
//	c.JSON(http.StatusOK, gin.H{"message": "This is Assets"})
//}

// @Summary Get finance data for a specific asset and timeframe
// @Description Returns financial data based on the selected asset and timeframe
// @Tags finance
// @Accept  json
// @Produce  json
// @Param asset path string true "Asset ID"
// @Param timeframe path string true "Timeframe ID"
// @Success 200 {object} models.FinanceData
// @Failure 404 {object} map[string]string
// @Router /finance/data/{asset}/{timeframe} [get]
//func (fc *FinanceController) handleData(c *gin.Context) {
//	c.JSON(http.StatusOK, gin.H{"message": "This is data"})
//}
