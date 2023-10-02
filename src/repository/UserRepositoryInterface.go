package repository

import (
	"database/sql"
	errApi "github.com/arilsonsantos/crud-mvc-go.git/src/config/error"
	"github.com/arilsonsantos/crud-mvc-go.git/src/model/domain"
)

type UserRepositoryInterface interface {
	Create(domainInterface domain.UserDomainInterface) (domain.UserDomainInterface, *errApi.ErrApi)
}

type userRepositoryInterface struct {
	sql *sql.DB
}

func NewUserRepository(sql *sql.DB) UserRepositoryInterface {
	return &userRepositoryInterface{sql}
}
