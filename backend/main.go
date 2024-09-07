package main

import (
	"healthcare-app/config" // Import the config package
	"healthcare-app/routes" // Import the routes package
)

func main() {
	// Initialize the database
	config.InitDatabase()

	// Initialize the router
	router := routes.SetupRouter()

	// Start the server
	router.Run(":8080")
}
