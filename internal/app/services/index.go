package app

import (
	"app/internal/domain/entities"
	"time"
)

type IUserService interface {
	Create(name, email, password, role string) (*entities.User, error)
}

type IAppointmentService interface {
	Schedule(barberID, customerID, offeringID string, startAt time.Time) (*entities.Appointment, error)
}
