package services

import (
	"bookstore/pkg/models"
	"bookstore/pkg/repositories"
	"errors"
)

type UserService struct {
	repository repositories.IUserRepository
}

func NewUserService(repo repositories.IUserRepository) *UserService {
	return &UserService{repository: repo}
}

func (s *UserService) GetAll() ([]models.User, error) {
	return s.repository.GetAll()
}

func (s *UserService) GetById(id string) (models.User, error) {
	return s.repository.GetById(id)
}

func (s *UserService) CreateUser(user models.User) (models.User, error) {
	return s.repository.CreateUser(user)
}

func (s *UserService) DeleteUser(id string) error {
	return s.repository.DeleteUser(id)
}
func (s *UserService) CheckUniqueName(name string) error {
	users, err := s.GetAll()
	if err != nil {
		return err
	}
	for _, user := range users {
		if user.Name == name {
			return errors.New("name already exists")
		}
	}
	return nil
}
