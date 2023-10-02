package controller

import (
	"github.com/arilsonsantos/crud-mvc-go.git/src/config/logger"
	"github.com/arilsonsantos/crud-mvc-go.git/src/controller/dto"
	"github.com/arilsonsantos/crud-mvc-go.git/src/model/domain"
	"github.com/arilsonsantos/crud-mvc-go.git/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	var requestDto dto.UserRequestDto

	if err := c.ShouldBind(&requestDto); err != nil {
		c.JSON(500, err)
	}

	userDomain := domain.NewUserDomain(
		requestDto.Name,
		requestDto.Password,
		requestDto.Email,
		requestDto.Age,
	)

	userDomainResult, err := uc.userService.CreateUser(userDomain)
	if err != nil {
		logger.Error("Error trying to create user.",
			err,
			zap.String("UserControllerInterface", "CreateUser"))
		println(err.DetailedMessage)
	}

	logger.Info("User created successfully.",
		zap.String("Local", "UserControllerInterface"),
		zap.String("Method", "CreateUser"))

	c.JSON(http.StatusCreated, view.ConvertUserDomainToUserDto(userDomainResult))

}
