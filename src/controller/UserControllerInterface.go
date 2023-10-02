package controller

import (
	"github.com/arilsonsantos/crud-mvc-go.git/src/service"
	"github.com/gin-gonic/gin"
)

type UserControllerInterface interface {
	CreateUser(c *gin.Context)
}

type userControllerInterface struct {
	userService service.UserServiceInterface
}

func NewUserControllerInterface(userService service.UserServiceInterface) UserControllerInterface {
	return &userControllerInterface{
		userService: userService,
	}

}
