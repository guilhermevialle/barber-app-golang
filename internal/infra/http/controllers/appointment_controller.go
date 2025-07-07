package controllers

import (
	app "app/internal/app/services"
	"app/internal/interfaces/dtos"
	"time"

	"github.com/gin-gonic/gin"
)

type AppointmentController struct {
	service app.IAppointmentService
}

var _ IAppointmentController = (*AppointmentController)(nil)

func NewAppointmentController(service app.IAppointmentService) *AppointmentController {
	return &AppointmentController{
		service: service,
	}
}

func (ac *AppointmentController) ScheduleAppointment(ctx *gin.Context) {
	var req dtos.ScheduleAppointmentDto

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	startAt, err := time.Parse(time.RFC3339, req.StartAt)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid date format. Use ISO8601 format."})
		return
	}

	appointment, err := ac.service.Schedule(req.BarberID, req.CustomerID, startAt)
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}

	ctx.JSON(200, appointment)
}
