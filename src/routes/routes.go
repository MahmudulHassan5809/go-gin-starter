package routes

import (
	"gin_starter/src/modules/auth"
	"gin_starter/src/modules/health"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterRoutes sets up routes for all modules
func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	api := r.Group("/api")
	{
		health.RegisterHealthRoutes(api.Group("/health-check"))
		auth.RegisterAuthRoutes(api.Group("/auth"), db)

	}
}
