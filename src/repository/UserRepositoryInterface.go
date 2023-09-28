package repository

import (
	"github.com/arilsonsantos/crud-mvc-go.git/src/model/domain"
	"google.golang.org/protobuf/types/known/apipb"
)

type UserRepositoryInterface interface {
	Create(domainInterface domain.UserDomainInterface) (domain.UserDomainInterface, *apipb.Api)
}

type userRepositoryInterface struct {
	connection string
}
