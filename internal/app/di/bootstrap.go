package di

import (
	app "app/internal/app/services"
	"app/internal/infra/http/controllers"
	imr "app/internal/infra/repositories/in_memory"
)

type DI struct {
	UserController        controllers.IUserController
	AppointmentController controllers.IAppointmentController
}

func NewDIBootstrap() *DI {
	// offering
	or := imr.NewOfferingRepository()

	// user
	ur := imr.NewUserRepository()
	us := app.NewUserService(ur)
	uc := controllers.NewUserController(us)

	// appointment
	ar := imr.NewAppointmentRepository()
	as := app.NewAppointmentService(ur, ar, or)
	ac := controllers.NewAppointmentController(as)

	return &DI{
		UserController:        uc,
		AppointmentController: ac,
	}
}
