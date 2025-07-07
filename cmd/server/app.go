package server

import (
	"app/internal/app/di"
	"app/internal/infra/http/routes"

	"github.com/gin-gonic/gin"
)

func NewApp() *gin.Engine {
	r := gin.Default()

	userController := di.NewUserDI()
	appointmentController := di.NewAppointmentDI()

	routes.RegisterUserRoutes(r, userController)
	routes.RegisterAppointmentRoutes(r, appointmentController)

	return r
}
