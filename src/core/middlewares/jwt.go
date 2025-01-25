package middlewares

import (
	"context"
	"fmt"
	"gin_starter/src/core/cache"
	"gin_starter/src/core/common"
	"gin_starter/src/core/errors"
	"gin_starter/src/core/security"
	"strings"

	"github.com/gin-gonic/gin"
)


type JwtMiddleware struct {
	CacheManager *cache.CacheManager
	JWTHandler security.JWTHandler
}

func NewJwtMiddleware(CacheManager *cache.CacheManager, JWTHandler security.JWTHandler) JwtMiddleware {
	return JwtMiddleware{CacheManager: CacheManager, JWTHandler: JWTHandler}
}

func (j *JwtMiddleware) JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(401, gin.H{"error": "missing or invalid Authorization header"})
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		valid, claims, err := j.VerifyJWT(token)

		if err != nil || !valid {
			c.AbortWithStatusJSON(401, gin.H{"error": "invalid or expired token"})
			return
		}

		c.Set("user_id", claims.UserID)
		c.Next()
	}
}

func (j *JwtMiddleware) VerifyJWT(token string) (bool, *common.AccessTokenPayload, error) {
	parsedToken, err := j.JWTHandler.Decode(token)
	if err != nil {
		return false, nil, errors.UnauthorizedError("invalid or expired token")
	}
	claims := &common.AccessTokenPayload{
		UserID: parsedToken["user_id"].(string),
		Email:  parsedToken["email"].(string),
		Sub:    parsedToken["sub"].(string),
	}

	if claims.Sub != "ACCESS_TOKEN" {
		return false, nil, errors.UnauthorizedError("invalid token subject")
	}

	cacheKey := cache.UserAccessToken.Format(claims.UserID)
	cachedToken, err := j.CacheManager.Get(context.Background(), cacheKey)
	if err != nil || cachedToken != token {
		return false, nil, errors.UnauthorizedError("token mismatch or not found in cache")
	}

	userDataCacheKey := fmt.Sprintf("USER_DATA_%v", claims.UserID)
	userData, err := j.CacheManager.HGetAll(context.Background(), userDataCacheKey)
	if err != nil || userData == nil {
		return false, nil, errors.UnauthorizedError("user data not found in cache")
	}

	return true, claims, nil
}
