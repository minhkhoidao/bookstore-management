package controllers

import (
	"bookstore/pkg/models"
	"bookstore/pkg/services"
	"bookstore/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *Controller {
	return &Controller{userService: userService}
}

func (c *Controller) GetAll(ctx *gin.Context) {
	users, err := c.userService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}
	ctx.JSON(http.StatusOK, users)
}

func (c *Controller) GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := c.userService.GetById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	}
	ctx.JSON(http.StatusOK, user)
}

func (c *Controller) CreateUser(ctx *gin.Context) {
	var user models.User
	id, errors := utils.GenerateRandomID(16)

	if errors != nil {
		return
	}
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	user.ID = id

	if err := c.userService.CheckUniqueName(user.Name); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Name already exists"})
		return
	}
	createdUser, err := c.userService.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}
	ctx.JSON(http.StatusCreated, createdUser)
}
