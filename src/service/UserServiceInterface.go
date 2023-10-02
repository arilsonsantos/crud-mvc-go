package service

import (
	errApi "github.com/arilsonsantos/crud-mvc-go.git/src/config/error"
	"github.com/arilsonsantos/crud-mvc-go.git/src/model/domain"
	"github.com/arilsonsantos/crud-mvc-go.git/src/repository"
)

type UserServiceInterface interface {
	CreateUser(udi domain.UserDomainInterface) (domain.UserDomainInterface, *errApi.ErrApi)
}

type userServiceInterface struct {
	userRepository repository.UserRepositoryInterface
}

func NewUserServiceInterface(ur repository.UserRepositoryInterface) UserServiceInterface {
	return &userServiceInterface{
		userRepository: ur,
	}
}
