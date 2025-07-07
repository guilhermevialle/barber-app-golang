package in_memory_repository

import (
	"app/internal/domain/entities"
	"app/internal/infra/repositories"
)

type AppointmentRepository struct {
	appointments []*entities.Appointment
}

var _ repositories.IAppointmentRepository = (*AppointmentRepository)(nil)

func NewAppointmentRepository() *AppointmentRepository {
	return &AppointmentRepository{
		appointments: make([]*entities.Appointment, 0),
	}
}

func (ar *AppointmentRepository) Save(appointment *entities.Appointment) {
	ar.appointments = append(ar.appointments, appointment)
}
