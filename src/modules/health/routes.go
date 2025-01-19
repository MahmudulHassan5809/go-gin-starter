package health

import "github.com/gin-gonic/gin"

// RegisterHealthRoutes sets up health-check routes
func RegisterHealthRoutes(r *gin.RouterGroup) {
	r.GET("/", HealthCheckHandler)
}
