package controllers

import (
	"log"
	"net/http"
	"time"

	"healthcare-app/config"
	"healthcare-app/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// Define your JWT secret key (ensure it's securely managed)
var jwtSecret = []byte("your_jwt_secret_key")

// Signup handler
func Signup(c *gin.Context) {
	var request struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Log request data for debugging
	log.Printf("Signup request received: Name=%s, Email=%s", request.Name, request.Email)

	// Parse the request body
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("Error parsing signup request: %v", err) // Log the error
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v", err) // Log the error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Create a new user and save it to the database
	user := models.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: string(hashedPassword),
	}

	result := config.DB.Create(&user)
	if result.Error != nil {
		log.Printf("Error creating user in the database: %v", result.Error) // Log the error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Log successful user creation
	log.Printf("User created successfully: Email=%s", request.Email)

	// Return a successful response
	c.JSON(http.StatusOK, gin.H{"message": "Signup successful", "user": user})
}

// Login handler with JWT
func Login(c *gin.Context) {
	var request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Log request data for debugging
	log.Printf("Login request received: Email=%s", request.Email)

	// Parse the request body
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("Error parsing login request: %v", err) // Log the error
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Retrieve the user from the database
	var user models.User
	result := config.DB.Where("email = ?", request.Email).First(&user)
	if result.Error != nil {
		log.Printf("User not found: Email=%s", request.Email) // Log the error
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	// Compare the stored hashed password with the provided password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		log.Printf("Invalid password for user: Email=%s", request.Email) // Log the error
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	// Generate JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(), // Token expiration: 72 hours
	})

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		log.Printf("Error generating token: %v", err) // Log the error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Log successful login
	log.Printf("User logged in successfully: Email=%s", request.Email)

	// Return the token
	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": tokenString})
}
