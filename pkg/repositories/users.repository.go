package repositories

import (
	"bookstore/pkg/models"
	"errors"
)

type UserRepository struct {
	users []models.User
}

type IUserRepository interface {
	GetAll() ([]models.User, error)
	GetById(id string) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	DeleteUser(id string) error
}

func (r *UserRepository) GetAll() ([]models.User, error) {
	return r.users, nil
}

func (r *UserRepository) GetById(id string) (models.User, error) {
	for _, user := range r.users {
		if user.ID == id {
			return user, nil
		}
	}
	return models.User{}, errors.New("user not found")
}

func (r *UserRepository) CreateUser(user models.User) (models.User, error) {
	r.users = append(r.users, user)
	return user, nil
}

func (r *UserRepository) DeleteUser(id string) error {
	for i, user := range r.users {
		if user.ID == id {
			r.users = append(r.users[:i], r.users[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}
