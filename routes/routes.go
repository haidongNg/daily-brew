package routes

import (
	"daily-brew/controllers"
	"daily-brew/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	protected := r.Group("/")
	protected.Use(middlewares.AuthMiddleware())
	{
		protected.POST("/refresh", controllers.RefreshToken)
	}
	return r
}
