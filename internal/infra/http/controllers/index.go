package controllers

import "github.com/gin-gonic/gin"

type IUserController interface {
	CreateUser(ctx *gin.Context)
}

type IAppointmentController interface {
	ScheduleAppointment(ctx *gin.Context)
}
