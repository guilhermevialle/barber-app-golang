package app

import (
	"app/internal/domain/entities"
	"app/internal/infra/repositories"
	"time"
)

type AppointmentService struct {
	repo repositories.IAppointmentRepository
}

var _ IAppointmentService = (*AppointmentService)(nil)

func NewAppointmentService(repo repositories.IAppointmentRepository) *AppointmentService {
	return &AppointmentService{
		repo: repo,
	}
}

func (as *AppointmentService) Schedule(BarberID, CustomerID string, startAt time.Time) (*entities.Appointment, error) {
	appointment, err := entities.NewAppointment(BarberID, CustomerID, startAt)
	if err != nil {
		return nil, err
	}

	as.repo.Schedule(appointment)

	return appointment, nil
}
