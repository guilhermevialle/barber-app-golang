package entities

import (
	"errors"
	"time"

	nanoid "github.com/matoous/go-nanoid/v2"
)

type Appointment struct {
	ID         string
	BarberID   string
	CustomerID string
	OfferingID string
	StartAt    time.Time
	Duration   time.Duration // minutes
}

func NewAppointment(barberID, customerID, offeringID string, startAt time.Time, duration time.Duration) (*Appointment, error) {
	id, err := nanoid.New(21)
	if err != nil {
		return nil, errors.New("failed to generate appointment id")
	}

	return &Appointment{
		ID:         id,
		BarberID:   barberID,
		CustomerID: customerID,
		OfferingID: offeringID,
		StartAt:    startAt,
		Duration:   duration,
	}, nil
}

func (a *Appointment) EndAt() time.Time {
	return a.StartAt.Add(a.Duration * time.Minute)
}

func (a *Appointment) DurationInMinutes() int {
	return int(a.Duration.Minutes())
}
