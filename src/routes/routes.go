package routes

import (
	"gin_starter/src/modules/auth"
	"gin_starter/src/modules/health"
	"gin_starter/src/modules/users"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	api := r.Group("/api")
	{
		health.RegisterHealthRoutes(api.Group("/health-check"))
		auth.RegisterAuthRoutes(api.Group("/auth"), db)
		users.RegisterUserRoutes(api.Group("/users"), db)

	}
}
