package entities

import (
	"errors"

	nanoid "github.com/matoous/go-nanoid/v2"
)

type User struct {
	ID       string
	Name     string
	Email    string
	Password string
	Role     string
}

func NewUser(name, email, password, role string) (*User, error) {
	id, err := nanoid.New(21)
	if err != nil {
		return nil, errors.New("failed to generate user id")
	}

	return &User{
		ID:       id,
		Name:     name,
		Email:    email,
		Password: password,
		Role:     role,
	}, nil
}
