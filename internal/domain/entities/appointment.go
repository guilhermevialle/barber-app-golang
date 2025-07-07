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
}

func NewAppointment(BarberID, CustomerID, OfferingID string, startAt time.Time) (*Appointment, error) {
	id, err := nanoid.New(21)
	if err != nil {
		return nil, errors.New("failed to generate appointment id")
	}

	return &Appointment{
		ID:         id,
		BarberID:   BarberID,
		CustomerID: CustomerID,
		OfferingID: OfferingID,
		StartAt:    startAt,
	}, nil
}
