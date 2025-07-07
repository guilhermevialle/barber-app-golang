package repositories

import (
	"app/internal/domain/entities"
)

type IUserRepository interface {
	Save(user *entities.User)
	FindById(id string) *entities.User
	FindByEmail(email string) *entities.User
	FindCustomerByID(customerID string) *entities.User
	FindBarberByID(barberID string) *entities.User
	FindByRoleAndID(role, id string) *entities.User
}

type IAppointmentRepository interface {
	Save(appointment *entities.Appointment)
}

type IOfferingRepository interface {
	Save(offering *entities.Offering)
	FindById(id string) *entities.Offering
}
