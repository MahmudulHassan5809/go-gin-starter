package health

import (
	"gin_starter/src/core/cache"

	"github.com/gin-gonic/gin"
)

// RegisterHealthRoutes sets up health-check routes
func RegisterHealthRoutes(r *gin.RouterGroup,) {
	CacheManager := cache.GetCacheManager()
	r.GET("/", cache.CacheResponse(CacheManager, "HEALTH", 60, HealthCheckHandler))
}