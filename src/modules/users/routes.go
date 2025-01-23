package users

import (
	"gin_starter/src/core/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUserRoutes(r *gin.RouterGroup, db *gorm.DB) {

	queryService := services.NewQueryService()
	
	repo := NewUserRepository(db)
	service := NewUserService(repo)
	handler := NewUserHandler(service, queryService)
	
	
	r.GET("/", handler.UserList)
}
