package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"safehouse-main-back/src/internal/database"
	"safehouse-main-back/src/internal/models"
	"safehouse-main-back/src/internal/service"
)

type AboutController struct {
	db                    database.Database
	personalReviewService service.PersonalReviewService
}

func (ac *AboutController) RegisterRoutes(router *gin.Engine) {
	router.GET("/about/contact", ac.handleContactRequest)
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

// @Summary Get random reviews from random people
// @Description Get random reviews from random people, for the carousel component in the about section
// @Tags about
// @Accept  json
// @Produce  json
// @Success 200 {object} []*models.PersonalReviewsCarouselDTO
// @Failure 404 {object} map[string]string
// @Router /about/reviews/carousel [get]
func (ac *AboutController) handleReviewsCarouselRequest(c *gin.Context) {
	reviewsCarouselDTOs := ac.personalReviewService.GetAllReviews()
	c.JSON(http.StatusOK, reviewsCarouselDTOs)
}
