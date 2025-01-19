package routes

import (
	"gin_starter/src/modules/auth"
	"gin_starter/src/modules/health"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

// RegisterRoutes sets up routes for all modules
func RegisterRoutes(r *gin.Engine, db *pgxpool.Pool) {
	api := r.Group("/api")
	{
		health.RegisterHealthRoutes(api.Group("/health-check"))
		auth.RegisterAuthRoutes(api.Group("/auth"), db)

	}
}
