package cache

import (
	"bytes"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)


func CacheResponse(cacheManager *CacheManager, cacheKey string, ttl int, handlerFunc gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		cacheKey := cacheKey
		ctx := c.Request.Context()
		cachedData, err := cacheManager.Get(ctx, cacheKey)
		if err == nil && cachedData != "" {
			c.Data(http.StatusOK, "application/json", []byte(cachedData))
			return
		}
		writer := &responseCacheWriter{ResponseWriter: c.Writer, body: &bytes.Buffer{}}
		c.Writer = writer
		handlerFunc(c)
		if c.Writer.Status() == http.StatusOK {
			if err := cacheManager.Set(ctx, cacheKey, writer.body.String(), ttl); err != nil {
				log.Printf("Error caching response: %v", err)
			}
		}
	}
}

type responseCacheWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *responseCacheWriter) Write(data []byte) (int, error) {
	w.body.Write(data)
	return w.ResponseWriter.Write(data)
}
