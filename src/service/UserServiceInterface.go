package service

import (
	errApi "github.com/arilsonsantos/crud-mvc-go.git/src/config/error"
	"github.com/arilsonsantos/crud-mvc-go.git/src/model/domain"
)

type UserServiceInterface interface {
	CreateUser(udi domain.UserDomainInterface) (domain.UserDomainInterface, *errApi.ErrApi)
}

type userService struct {
}
