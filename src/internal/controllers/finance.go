package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type FinanceController struct {
	db *gorm.DB
}

func (fc *FinanceController) RegisterRoutes(router *gin.Engine) {
	router.GET("/finance/intro", fc.handleIntro)
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
