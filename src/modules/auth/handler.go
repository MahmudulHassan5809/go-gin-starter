package auth

import (
	"gin_starter/src/core/errors"
	"gin_starter/src/core/helpers"
	"gin_starter/src/modules/users"
	"net/http"

	"github.com/gin-gonic/gin"
)


type AuthHandler struct {
	authService AuthService
}

func NewAuthHandler(service AuthService) *AuthHandler {
	return &AuthHandler{authService: service}
}

func (h *AuthHandler) LoginHandler(ctx *gin.Context) {
	var loginRequest LoginRequest
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.Error(errors.BadRequestError(err.Error()))
		return
	}

	tokens, err := h.authService.LoginUser(&loginRequest)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, helpers.BuildSuccessResponse(
		tokens,
	))
}

func (h *AuthHandler) Register(ctx *gin.Context) {
	var createdUser users.CreateUserRequest
	if err := ctx.ShouldBindJSON(&createdUser); err != nil {
		ctx.Error(errors.BadRequestError(err.Error()))
		return
	}

	err := h.authService.RegisterUser(&createdUser)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, helpers.BuildSuccessResponse(
		gin.H{"status": "success"},
	))
}