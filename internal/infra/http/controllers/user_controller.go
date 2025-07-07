package controllers

import (
	app "app/internal/app/services"
	"app/internal/interfaces/dtos"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service app.IUserService
}

var _ IUserController = (*UserController)(nil)

func NewUserController(service app.IUserService) *UserController {
	return &UserController{
		service: service,
	}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	var req dtos.CreateUserDto

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	user, err := uc.service.Create(req.Name, req.Email, req.Password, req.Role)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, user)
}
