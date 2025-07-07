package routes

import (
	"app/internal/infra/http/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine, controller controllers.IUserController) {
	router.POST("/user/register", controller.CreateUser)
}
