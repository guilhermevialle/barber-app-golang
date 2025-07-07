package entities

import (
	"errors"

	nanoid "github.com/matoous/go-nanoid/v2"
)

type Offering struct {
	ID    string
	Name  string
	Price int
}

func NewOffering(name string, price int) (*Offering, error) {
	id, err := nanoid.New(21)
	if err != nil {
		return nil, errors.New("failed to generate offering id")
	}

	return &Offering{
		ID:    id,
		Name:  name,
		Price: price,
	}, nil
}
