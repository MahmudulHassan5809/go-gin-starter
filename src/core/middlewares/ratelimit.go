package middlewares

import (
	"context"
	"fmt"
	"gin_starter/src/core/ratelimiter"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func RateLimit(limit int, window time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		clientIP := c.ClientIP()
		key := "rate_limit:" + clientIP
		limiter := ratelimiter.NewRateLimiter(limit, window)
		allowed, err := limiter.Allow(ctx, key)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			c.Abort()
			return
		}

		if !allowed {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded"})
			c.Abort()
			return
		}
		c.Next()
	}
}
