package view

import (
	"github.com/arilsonsantos/crud-mvc-go.git/src/controller/dto"
	"github.com/arilsonsantos/crud-mvc-go.git/src/model/domain"
)

func ConvertUserDomainToUserDto(
	userDomain domain.UserDomainInterface,
) dto.UserResponseDto {
	return dto.UserResponseDto{
		Id:    userDomain.GetID(),
		Name:  userDomain.GetName(),
		Email: userDomain.GetEmail(),
		Age:   userDomain.GetAge(),
	}
}
