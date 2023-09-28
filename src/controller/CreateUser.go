package controller

import (
	errApi "github.com/arilsonsantos/crud-mvc-go.git/src/config/error"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	err := errApi.NewBadRequest("Bateu aqui")
	if err != nil {
		c.JSON(err.HttpCode, err)
		return
	}
}
