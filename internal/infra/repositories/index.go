package repositories

import (
	"app/internal/domain/entities"
)

type IUserRepository interface {
	Save(user *entities.User)
	FindById(id string) *entities.User
	FindByEmail(email string) *entities.User
}

type IAppointmentRepository interface {
	Schedule(appointment *entities.Appointment)
}
