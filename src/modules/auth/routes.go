package auth

import (
	"gin_starter/src/core/cache"
	"gin_starter/src/core/security"
	"gin_starter/src/core/services"
	"gin_starter/src/modules/users"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func RegisterAuthRoutes(r *gin.RouterGroup, db *gorm.DB) {
	passwordHandler := security.PasswordHandler{}
	JwtHandler := security.JWTHandler{}
	validatorService := services.NewValidatorService()
	CacheManager := cache.GetCacheManager()
	
	repo := users.NewUserRepository(db)
	service := NewAuthService(repo, passwordHandler, JwtHandler, CacheManager)
	handler := NewAuthHandler(service, validatorService)
	
	
	r.POST("/register/", handler.Register)
	r.POST("/login/", handler.LoginHandler)
}
