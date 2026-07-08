package routes

import (
	"backend-test-golang/handlers"
	"backend-test-golang/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	users := r.Group("/users")
	users.Use(middleware.LogMiddleware())
	users.Use(middleware.AuthMiddleware())
	{
		users.GET("", handlers.GetAllUsers)
	}
}
