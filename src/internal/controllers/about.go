package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"safehouse-main-back/src/internal/database"
	"safehouse-main-back/src/internal/models"
)

type AboutController struct {
	db database.Database
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
	aboutIntroMessage := "Hey! I’m David—great to see you here!\n " +
		"I built a simple portfolio page a while back but kept postponing it—well, it’s finally live!\n " +
		"I’m a Senior Software Engineer specializing in backend. " +
		"I have a LOT of interests and tend to explore them fully, so expect to see some interesting projects here!\n " +
		"You’ll also find my contact info and a peek into my passions. Enjoy exploring!"
	c.JSON(http.StatusOK, gin.H{"message": aboutIntroMessage})
}

// @Summary Get introduction about the app
// @Description Get an introduction message for the about section
// @Tags about
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]string
// @Router /about/contact-text [get]
func (ac *AboutController) handleContactText(c *gin.Context) {
	contactText := "This is my contact information. Don't hesitate in dropping a message if you want to connect!"
	c.JSON(http.StatusOK, gin.H{"message": contactText})
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
	contact, _ := ac.db.GetContact()
	contactsDTO := models.ToContactsDTO(contact)
	if contact == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contact not found"})
		return
	}
	c.JSON(http.StatusOK, contactsDTO)
}
