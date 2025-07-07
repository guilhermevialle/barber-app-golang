package in_memory_repository

import "app/internal/domain/entities"

type UserRepository struct {
	users []*entities.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make([]*entities.User, 0),
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
