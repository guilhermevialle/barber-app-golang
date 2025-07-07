package server

import (
	"app/internal/app/di"
	"app/internal/infra/http/routes"

	"github.com/gin-gonic/gin"
)

func NewApp() *gin.Engine {
	r := gin.Default()
	di := di.NewDIBootstrap()

	routes.RegisterUserRoutes(r, di.UserController)
	routes.RegisterAppointmentRoutes(r, di.AppointmentController)

	return r
}
