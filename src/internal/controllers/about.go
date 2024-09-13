package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"safehouse-main-back/src/internal/models"
)

type AboutController struct {
	db *gorm.DB
}

func (ac *AboutController) RegisterRoutes(router *gin.Engine) {
	router.GET("/about/intro", ac.handleIntro)
	router.GET("/about/contact", ac.getContactRequest)
	router.GET("/about/contact-text", ac.handleContactText)
}

func (ac *AboutController) handleIntro(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "This is the About Intro screen."})
}

func (ac *AboutController) handleContactText(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "This is the Contacts screen."})
}

func (ac *AboutController) getContactRequest(c *gin.Context) {
	contact := ac.getContact()
	if contact == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contact not found"})
		return
	}
	c.JSON(http.StatusOK, contact)
}

func (ac *AboutController) getContact() *models.Contacts {
	var contact models.Contacts
	if err := ac.db.Where("active = ?", true).First(&contact).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		panic(err)
	}
	return &contact
}
