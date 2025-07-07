package di

import (
	app "app/internal/app/services"
	"app/internal/infra/http/controllers"
	imr "app/internal/infra/repositories/in_memory"
)

func NewUserDI() controllers.IUserController {
	repo := imr.NewUserRepository()
	service := app.NewUserService(repo)
	controller := controllers.NewUserController(service)

	return controller
}
