package in_memory_repository

import (
	"app/internal/domain/entities"
	"app/internal/infra/repositories"
)

type UserRepository struct {
	users []*entities.User
}

var _ repositories.IUserRepository = (*UserRepository)(nil)

func NewUserRepository() *UserRepository {
	testing_customer := &entities.User{
		ID:       "customer-1",
		Name:     "John Doe",
		Email:    "john@doe.com",
		Password: "123456",
		Role:     "customer",
	}

	testing_barber := &entities.User{
		ID:       "barber-1",
		Name:     "Denis Doe",
		Email:    "denis@doe.com",
		Password: "123456",
		Role:     "barber",
	}

	return &UserRepository{
		users: []*entities.User{testing_customer, testing_barber},
	}
}

func (ur *UserRepository) Save(user *entities.User) {
	ur.users = append(ur.users, user)
}

func (ur *UserRepository) FindById(id string) *entities.User {
	for _, user := range ur.users {
		if user.ID == id {
			return user
		}
	}
	return nil
}

func (ur *UserRepository) FindByEmail(email string) *entities.User {
	for _, user := range ur.users {
		if user.Email == email {
			return user
		}
	}

	return nil
}

func (ur *UserRepository) FindCustomerByID(customerID string) *entities.User {
	for _, user := range ur.users {
		if user.ID == customerID && user.Role == "customer" {
			return user
		}
	}

	return nil
}

func (ur *UserRepository) FindBarberByID(barberID string) *entities.User {
	for _, user := range ur.users {
		if user.ID == barberID && user.Role == "barber" {
			return user
		}
	}

	return nil
}

func (ur *UserRepository) FindByRoleAndID(role, id string) *entities.User {
	for _, user := range ur.users {
		if user.ID == id && user.Role == role {
			return user
		}
	}

	return nil
}
