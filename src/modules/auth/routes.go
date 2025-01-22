package auth

import (
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
	
	repo := users.NewUserRepository(db)
	service := NewAuthService(repo, passwordHandler, JwtHandler)
	handler := NewAuthHandler(service, validatorService)
	
	
	r.POST("/register/", handler.Register)
	r.POST("/login/", handler.LoginHandler)
}
