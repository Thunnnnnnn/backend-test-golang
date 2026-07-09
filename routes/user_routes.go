package routes

import (
	"backend-test-golang/handlers"
	"backend-test-golang/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	users := r.Group("/users")
	users.Use(middleware.LogMiddleware())
	{
		users.POST("", handlers.CreateUser)
	}
	users.Use(middleware.AuthMiddleware())
	{
		users.GET("", handlers.GetAllUsers)
		users.GET("/:id", handlers.GetUserByID)
		users.PUT("/:id", handlers.UpdateUser)
		users.DELETE("/:id", handlers.DeleteUser)
	}
}
