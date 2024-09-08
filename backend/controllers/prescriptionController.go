package controllers

import (
	"github.com/gin-gonic/gin"
	"healthcare-app/config"
	"healthcare-app/models"
	"net/http"
)

// CreatePrescription creates a new prescription for the user
func CreatePrescription(c *gin.Context) {
	var prescription models.Prescription

	// Bind the prescription details from the request
	if err := c.ShouldBindJSON(&prescription); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Save the prescription to the database
	if err := config.DB.Create(&prescription).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create prescription"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Prescription created successfully"})
}

// GetPrescriptions retrieves the user's prescriptions
func GetPrescriptions(c *gin.Context) {
	var prescriptions []models.Prescription
	email := c.GetString("email")

	// Retrieve prescriptions from the database using the user's email
	if err := config.DB.Where("email = ?", email).Find(&prescriptions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve prescriptions"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"prescriptions": prescriptions})
}
