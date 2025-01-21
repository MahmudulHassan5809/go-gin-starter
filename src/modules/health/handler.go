package health

import (
	"gin_starter/src/core/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)


func HealthCheckHandler(ctx *gin.Context) {
	response := helpers.BuildSuccessResponse(
		gin.H{
			"app": "Go Gin Starter",
			"ver": "1.0.0",
		},
		helpers.WithMessage("Health check successful"),
	)
	ctx.JSON(http.StatusOK, response)
}
