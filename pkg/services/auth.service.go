package services

import (
	"bookstore/pkg/models"
	"bookstore/pkg/repositories"
)

type AuthService struct {
	repository repositories.AuthRepository
}

func NewAuthService(repo repositories.AuthRepository) *AuthService {
	return &AuthService{repository: repo}
}

func (s *AuthService) Signin(auth models.Auth) (string, error) {
	return s.repository.Signin(auth)
}

func (s *AuthService) Signup(auth models.Auth) (string, error) {
	return s.repository.Signup(auth)
}
