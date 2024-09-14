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

func (fc *FinanceController) handleIntro(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "This is the Finance Intro screen."})
}

func (fc *FinanceController) handleNews(c *gin.Context) {
	service.GetNewsByGenre(c, models.NewsGenreTech, fc.db)
}

func (fc *FinanceController) handleTopicOfTheSeason(c *gin.Context) {
	service.GetTopicOfTheSeasonByGenre(c, models.NewsGenreTech, fc.db)
}

//func (fc *FinanceController) handleTimeframes(c *gin.Context) {
//	c.JSON(http.StatusOK, gin.H{"message": "This is Timeframes"})
//}
//
//func (fc *FinanceController) handleAssets(c *gin.Context) {
//	c.JSON(http.StatusOK, gin.H{"message": "This is Assets"})
//}
//
//func (fc *FinanceController) handleData(c *gin.Context) {
//	c.JSON(http.StatusOK, gin.H{"message": "This is data"})
//}
