package routes

import (
	"backend-test-golang/handlers"
	"backend-test-golang/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	auth.Use(middleware.LogMiddleware())
	{
		auth.POST("/login", handlers.Login)
		auth.POST("/register", handlers.CreateUser)
	}
}
