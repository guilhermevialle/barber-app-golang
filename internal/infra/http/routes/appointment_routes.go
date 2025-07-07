package routes

import (
	"app/internal/infra/http/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterAppointmentRoutes(router *gin.Engine, controller controllers.IAppointmentController) {
	router.POST("/appointment/schedule", controller.ScheduleAppointment)
}
