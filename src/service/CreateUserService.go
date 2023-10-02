package service

import (
	errApi "github.com/arilsonsantos/crud-mvc-go.git/src/config/error"
	"github.com/arilsonsantos/crud-mvc-go.git/src/model/domain"
)

func (us *userServiceInterface) CreateUser(userDomain domain.UserDomainInterface) (domain.UserDomainInterface, *errApi.ErrApi) {
	if userDomain.GetAge() > 100 {
		err := errApi.NewErrApi(
			400,
			"CRUD_GO-400",
			"Bad Request",
			"Idade maior que 100.",
		)
		return nil, err
	}

	userDomain.EncryptPassword()
	userDomain, err := us.userRepository.Create(userDomain)

	if err != nil {
		return nil, err
	}

	return userDomain, nil
}
