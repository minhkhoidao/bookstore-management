package container

import "bookstore/pkg/services"

type AuthContainer struct {
	authService *services.AuthService
}
