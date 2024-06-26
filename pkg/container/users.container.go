package container

import (
	"bookstore/pkg/repositories"
	"bookstore/pkg/services"
)

func (c *Container) UserRepository() repositories.IUserRepository {
	// You could initialize dependencies here if needed
	return &repositories.UserRepository{}
}

// UserService retrieves the UserService with its dependencies
func (c *Container) UserService() *services.UserService {
	return services.NewUserService(c.UserRepository())
}
