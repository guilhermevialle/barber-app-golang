package app

import (
	"app/internal/domain/entities"
	"app/internal/infra/repositories"
	"errors"
	"time"
)

type AppointmentService struct {
	userRepo        repositories.IUserRepository
	appointmentRepo repositories.IAppointmentRepository
	offeringRepo    repositories.IOfferingRepository
}

// implements IAppointmentService
var _ IAppointmentService = (*AppointmentService)(nil)

func NewAppointmentService(
	userRepo repositories.IUserRepository,
	appointmentRepo repositories.IAppointmentRepository,
	offeringRepo repositories.IOfferingRepository,
) *AppointmentService {
	return &AppointmentService{
		userRepo:        userRepo,
		appointmentRepo: appointmentRepo,
		offeringRepo:    offeringRepo,
	}
}

func (as *AppointmentService) Schedule(barberID, customerID, offeringID string, startAt time.Time) (*entities.Appointment, error) {
	offeringExists := as.offeringRepo.FindById(offeringID)
	if offeringExists == nil {
		return nil, errors.New("offering does not exist")
	}

	barberExists := as.userRepo.FindByRoleAndID("barber", barberID)
	if barberExists == nil {
		return nil, errors.New("barber does not exist")
	}

	customerExists := as.userRepo.FindByRoleAndID("customer", customerID)
	if customerExists == nil {
		return nil, errors.New("customer does not exist")
	}

	appointment, err := entities.NewAppointment(barberID, customerID, offeringID, startAt, offeringExists.Duration)
	if err != nil {
		return nil, err
	}

	as.appointmentRepo.Save(appointment)

	return appointment, nil
}
