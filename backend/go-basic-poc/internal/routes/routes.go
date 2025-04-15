package routes

import (
	api "go-basic-poc/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) { // Use gin.Engine
	r.GET("/ping", api.Ping)
	// r.GET("/users/:id", handlers.User) // Gin uses :id for path parameters
}
