package main

import (
	"bookstore/pkg/container"
	"bookstore/pkg/routes"
	"fmt"
)

func InitContainer() {
	initContainer := container.NewContainer()
	userService := initContainer.UserService()

	router := routes.UserRoutes(userService)

	router.Run(":8080")
}

func main() {
	fmt.Println("Hello, World!")
	InitContainer()

}
