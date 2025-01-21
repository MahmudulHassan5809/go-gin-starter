package auth

import (
	"gin_starter/src/core/security"
	"gin_starter/src/modules/users"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)


func RegisterAuthRoutes(r *gin.RouterGroup, db *pgxpool.Pool) {
	passwordHandler := security.PasswordHandler{}
	JwtHandler := security.JWTHandler{}
	
	repo := users.NewUserRepository(db)
	service := NewAuthService(repo, passwordHandler, JwtHandler)
	handler := NewAuthHandler(service)
	
	
	r.POST("/register/", handler.Register)
	r.POST("/login/", handler.LoginHandler)
}
