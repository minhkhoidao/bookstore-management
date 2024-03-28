package repositories

import "bookstore/pkg/models"

type AuthRepository struct{}

type IAuthRepository interface {
	Signin(auth models.Auth) (string, error)
	Signup(auth models.Auth) (string, error)
}

func (r *AuthRepository) Signin(auth models.Auth) (string, error) {
	return "", nil
}

func (r *AuthRepository) Signup(auth models.Auth) (string, error) {
	return "", nil
}
