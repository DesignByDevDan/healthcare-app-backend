package routes

import (
	"healthcare-app/controllers"
	"healthcare-app/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Public routes
	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)

	// Protected routes (require a valid JWT token)
	protected := router.Group("/api")
	protected.Use(middleware.JWTAuthMiddleware()) // Apply JWT middleware

	// Add the protected resource route
	protected.GET("/protected-resource", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Access granted to protected resource!"})
	})

	return router
}
