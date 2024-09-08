package controllers

import (
	"github.com/gin-gonic/gin"
	"healthcare-app/config"
	"healthcare-app/models"
	"net/http"
)

// GetUserProfile retrieves the user's profile information
func GetUserProfile(c *gin.Context) {
	var user models.User
	email := c.GetString("email")

	// Retrieve the user profile from the database using the email
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"profile": user})
}

// UpdateUserProfile updates the user's profile information
func UpdateUserProfile(c *gin.Context) {
	var user models.User
	email := c.GetString("email")

	// Retrieve the user profile from the database
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		return
	}

	// Bind new profile data
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Save updated profile to the database
	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
}
