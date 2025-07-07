package app

import (
	"app/internal/domain/entities"
	"app/internal/infra/repositories"
	"errors"
)

type UserService struct {
	repo repositories.IUserRepository
}

// implements IUserService
var _ IUserService = (*UserService)(nil)

func NewUserService(repo repositories.IUserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (us *UserService) Create(name, email, password, role string) (*entities.User, error) {
	userExists := us.repo.FindByEmail(email)

	if userExists != nil {
		return nil, errors.New("user already exists")
	}

	user, err := entities.NewUser(name, email, password, role)
	if err != nil {
		return nil, err
	}

	us.repo.Save(user)

	return user, nil
}
