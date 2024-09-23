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
	router.GET("/about/contact", ac.handleContactRequest)
	router.GET("/about/contact-text", ac.handleContactText)
}

// @Summary Get introduction about the app
// @Description Get an introduction message for the about section
// @Tags about
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]string
// @Router /about/intro [get]
func (ac *AboutController) handleIntro(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "This is the About Intro screen."})
}

// @Summary Get introduction about the app
// @Description Get an introduction message for the about section
// @Tags about
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]string
// @Router /about/contact-text [get]
func (ac *AboutController) handleContactText(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "This is the Contacts screen."})
}

// @Summary Get contact information
// @Description Get contact information from the database
// @Tags about
// @Accept  json
// @Produce  json
// @Success 200 {object} models.ContactsDTO
// @Failure 404 {object} map[string]string
// @Router /about/contact [get]
func (ac *AboutController) handleContactRequest(c *gin.Context) {
	contact := ac.getContact()
	if contact == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contact not found"})
		return
	}
	c.JSON(http.StatusOK, contact)
}

func (ac *AboutController) getContact() *models.ContactsDTO {
	var contact models.Contacts
	if err := ac.db.Where("active = ?", true).First(&contact).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		panic(err)
	}
	contactsDTO := models.ToContactsDTO(&contact)
	return contactsDTO
}
