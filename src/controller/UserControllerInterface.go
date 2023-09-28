package controller

import "github.com/gin-gonic/gin"

type UserControllerInterface interface {
	CreateUser(c *gin.Context)
	FindUserById(c *gin.Context)
	FindUserByEmail(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}
