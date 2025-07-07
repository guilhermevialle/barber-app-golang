package di

import (
	app "app/internal/app/services"
	"app/internal/infra/http/controllers"
	imr "app/internal/infra/repositories/in_memory"
)

func NewAppointmentDI() controllers.IAppointmentController {
	repo := imr.NewAppointmentRepository()
	service := app.NewAppointmentService(repo)
	controller := controllers.NewAppointmentController(service)

	return controller
}
