package routes

import (
	"github.com/gin-gonic/gin"
	"healthcare-app/controllers"
	"healthcare-app/middleware"
)

func SetupRouter(router *gin.Engine) {
	// Public routes
	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)

	// Protected routes (require valid JWT token)
	protected := router.Group("/api")
	protected.Use(middleware.JWTAuthMiddleware()) // Apply JWT Middleware

	protected.GET("/protected-resource", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Access granted to protected resource!"})
	})
}
