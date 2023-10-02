package entity

import "github.com/arilsonsantos/crud-mvc-go.git/src/model/domain"

type UserEntity struct {
	id    string
	email string
	senha string
	nome  string
	idade int8
}

func UserDomainToEntity(userDomain domain.UserDomainInterface) *UserEntity {
	return &UserEntity{
		id:    userDomain.GetID(),
		email: userDomain.GetEmail(),
		nome:  userDomain.GetName(),
		senha: userDomain.GetPassword(),
		idade: userDomain.GetAge(),
	}
}

func UserEntityToDomain(userEntity UserEntity) domain.UserDomainInterface {
	userDomain := domain.NewUserDomain(
		userEntity.nome,
		userEntity.email,
		userEntity.senha,
		userEntity.idade,
	)
	userDomain.SetID(userEntity.id)

	return userDomain
}
