package main

import (
	"healthcare-app/config" // Import the config package for DB initialization
	"healthcare-app/routes" // Import the routes package for setting up routes

	"github.com/gin-contrib/cors" // Import the CORS middleware package
	"github.com/gin-gonic/gin"    // Import the Gin package for routing
)

func main() {
	// Initialize the database
	config.InitDatabase()

	// Initialize the router
	router := gin.Default()

	// Apply CORS middleware to allow cross-origin requests
	router.Use(cors.Default())

	// Set up the routes
	routes.SetupRouter()

	// Start the server on port 8080
	router.Run(":8080")
}
