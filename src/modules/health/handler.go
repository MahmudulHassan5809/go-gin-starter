package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheckHandler responds with the status of the application
func HealthCheckHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "healthy"})
}
