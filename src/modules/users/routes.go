package users

import (
	"gin_starter/src/core/cache"
	"gin_starter/src/core/middlewares"
	"gin_starter/src/core/security"
	"gin_starter/src/core/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUserRoutes(r *gin.RouterGroup, db *gorm.DB) {

	queryService := services.NewQueryService()
	
	repo := NewUserRepository(db)
	service := NewUserService(repo)
	handler := NewUserHandler(service, queryService)

	JwtHandler := security.JWTHandler{}
	CacheManager := cache.GetCacheManager()
	jwtMiddleware := middlewares.NewJwtMiddleware(CacheManager, JwtHandler)
	protectedRoutes := r.Group("/", jwtMiddleware.JwtMiddleware())
	
	protectedRoutes.GET("/", handler.UserList)
}
