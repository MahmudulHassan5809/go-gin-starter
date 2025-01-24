package health

import (
	"gin_starter/src/core/cache"
	"gin_starter/src/core/middlewares"
	"time"

	"github.com/gin-gonic/gin"
)


func RegisterHealthRoutes(r *gin.RouterGroup,) {
	CacheManager := cache.GetCacheManager()
	r.GET("/", middlewares.RateLimit(3, 10*time.Second), cache.CacheResponse(CacheManager, "HEALTH", 60, HealthCheckHandler))
}