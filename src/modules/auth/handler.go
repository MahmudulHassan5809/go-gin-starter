package auth

import (
	"gin_starter/src/core/errors"
	"gin_starter/src/core/helpers"
	"gin_starter/src/core/services"
	"gin_starter/src/modules/users"
	"net/http"

	"github.com/gin-gonic/gin"
)


type AuthHandler struct {
	authService AuthService
	validator *services.ValidatorService
}

func NewAuthHandler(service AuthService, validator *services.ValidatorService) *AuthHandler {
	return &AuthHandler{authService: service, validator: validator}
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
	var createdUser users.RegisterUserRequest
	if err := ctx.ShouldBindJSON(&createdUser); err != nil {
		ctx.Error(errors.BadRequestError(err.Error()))
		return
	}

	if err := h.validator.ValidateStruct(&createdUser); err != nil {
		validationErrors := helpers.ParseValidationErrors(err)
		ctx.Error(errors.BadRequestError("Validation failed").WithValidationErrors(validationErrors))
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