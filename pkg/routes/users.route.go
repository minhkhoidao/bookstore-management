package routes

import (
	"bookstore/pkg/controllers"
	"bookstore/pkg/services"

	"github.com/gin-gonic/gin"
)

func UserRoutes(userService *services.UserService) *gin.Engine {
	// Define the user routes here
	userController := controllers.NewUserController(userService)
	router := gin.Default()
	router.GET("/users", userController.GetAll)
	router.POST("/create", userController.CreateUser)
	return router
}
