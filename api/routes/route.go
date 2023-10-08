package routes

import (
	"cognixus/api/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Setup(db *sqlx.DB, gin *gin.Engine) {
	// publicRouter := gin.Group("api/v1")

	// All Public APIs
	// NewTodoRoute(db, publicRouter)

	protectedRouter := gin.Group("api/v1")
	// // Middleware to verify AccessToken
	protectedRouter.Use(middlewares.JwtAuthMiddleware())
	NewTodoRoute(db, protectedRouter)
}
