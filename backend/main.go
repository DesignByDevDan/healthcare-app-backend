package main

import (
	"github.com/gin-contrib/cors" // Import the CORS middleware
	"github.com/gin-gonic/gin"    // Import gin package
	"healthcare-app/config"       // Import the config package
	"healthcare-app/routes"       // Import the routes package
)

func main() {
	// Initialize the database
	config.InitDatabase()

	// Initialize the router
	router := gin.Default()

	// Apply CORS middleware
	router.Use(cors.Default())

	// Set up the routes
	routes.SetupRouter(router)

	// Start the server on port 8080
	router.Run(":8080")
}
