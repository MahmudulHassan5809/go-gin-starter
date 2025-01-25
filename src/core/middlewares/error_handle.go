package middlewares

import (
	"gin_starter/src/core/errors"
	"gin_starter/src/core/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandlingMiddleware() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        ctx.Next()

        // Handle errors after all middlewares and handlers
        if len(ctx.Errors) > 0 {
            err := ctx.Errors.Last()
            if customErr, ok := err.Err.(*errors.CustomError); ok {
                ctx.AbortWithStatusJSON(customErr.Status, helpers.BuildErrorResponse(
                    helpers.WithError(customErr),
                ))
            } else {
                ctx.AbortWithStatusJSON(http.StatusInternalServerError, helpers.BuildErrorResponse(
                    helpers.WithError(errors.InternalServerError("internal server error")),
                ))
            }
            return
        }
    }
}

