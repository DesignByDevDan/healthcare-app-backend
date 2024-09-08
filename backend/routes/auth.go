package routes

import (
	"healthcare-app/controllers"
	"healthcare-app/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Public routes can go here if you have any

	// Protected routes with JWT middleware
	protected := router.Group("/protected")
	protected.Use(middleware.JWTAuthMiddleware())
	{
		protected.GET("/profile", controllers.GetUserProfile)
		protected.PUT("/profile", controllers.UpdateUserProfile)
		protected.POST("/appointments", controllers.CreateAppointment)
		protected.GET("/appointments", controllers.GetAppointments)
		protected.POST("/prescriptions", controllers.CreatePrescription)
		protected.GET("/prescriptions", controllers.GetPrescriptions)
		protected.GET("/medical-history", controllers.GetMedicalHistory)
	}

	return router
}
