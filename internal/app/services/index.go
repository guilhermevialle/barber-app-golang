package app

import (
	"app/internal/domain/entities"
	"time"
)

type IUserService interface {
	Create(name, email, password, role string) (*entities.User, error)
}

type IAppointmentService interface {
	Schedule(BarberID, CustomerID, OfferingID string, startAt time.Time) (*entities.Appointment, error)
}
