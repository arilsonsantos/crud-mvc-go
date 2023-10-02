package routes

import (
	"github.com/arilsonsantos/crud-mvc-go.git/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, controller controller.UserControllerInterface) {
	const usersUserId = "/users/:userId"

	r.POST("/users", controller.CreateUser)
	r.GET(usersUserId)
	r.GET("/users/email/:userEmail")
	r.PUT(usersUserId)
	r.DELETE(usersUserId)
}
