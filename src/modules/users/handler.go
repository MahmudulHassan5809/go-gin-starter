package users

import (
	"gin_starter/src/core/errors"
	"gin_starter/src/core/helpers"
	"gin_starter/src/core/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService UserService
	queryService *services.QueryService
}

func NewUserHandler(userService UserService, queryService *services.QueryService) UserHandler {
	return UserHandler{userService: userService, queryService: queryService}
}

func (u *UserHandler) UserList(ctx *gin.Context) {
	filterOptions, err := u.queryService.ParseQueryParams(ctx, []string{"is_admin"}, []string{"username", "email"})
	if err != nil {
		ctx.Error(errors.BadRequestError(err.Error()))
		return
	}

	users, totalCount, err := u.userService.UserList(filterOptions, &User{})
	if err != nil {
		ctx.Error(errors.BadRequestError(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helpers.BuildSuccessResponse(
		users,
		helpers.WithMetaInfo(helpers.GeneratePaginationMeta(
			totalCount,
			filterOptions.Pagination.Page,
			filterOptions.Pagination.PageSize,
		)),
		
	))
}