package controllers

import (
	"github.com/gin-gonic/gin"
	"healthcare-app/config"
	"healthcare-app/models"
	"net/http"
)

// GetMedicalHistory retrieves the user's medical history
func GetMedicalHistory(c *gin.Context) {
	var medicalHistory []models.MedicalHistory
	email := c.GetString("email")

	// Retrieve medical history from the database using the user's email
	if err := config.DB.Where("email = ?", email).Find(&medicalHistory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve medical history"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"medicalHistory": medicalHistory})
}
