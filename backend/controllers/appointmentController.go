package controllers

import (
	"github.com/gin-gonic/gin"
	"healthcare-app/config"
	"healthcare-app/models"
	"net/http"
)

// CreateAppointment creates a new appointment for the user
func CreateAppointment(c *gin.Context) {
	var appointment models.Appointment

	// Bind the appointment details from the request
	if err := c.ShouldBindJSON(&appointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Save the appointment to the database
	if err := config.DB.Create(&appointment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create appointment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Appointment created successfully"})
}

// GetAppointments retrieves the user's appointments
func GetAppointments(c *gin.Context) {
	var appointments []models.Appointment
	email := c.GetString("email")

	// Retrieve appointments from the database using the user's email
	if err := config.DB.Where("email = ?", email).Find(&appointments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve appointments"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"appointments": appointments})
}
