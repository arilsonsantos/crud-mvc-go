package repository

import (
	errApi "github.com/arilsonsantos/crud-mvc-go.git/src/config/error"
	"github.com/arilsonsantos/crud-mvc-go.git/src/config/logger"
	"github.com/arilsonsantos/crud-mvc-go.git/src/model/domain"
	"github.com/arilsonsantos/crud-mvc-go.git/src/repository/entity"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (u userRepositoryInterface) Create(userDI domain.UserDomainInterface) (domain.UserDomainInterface, *errApi.ErrApi) {
	id, _ := uuid.NewUUID()
	userDI.SetID(id.String())
	userEntity := entity.UserDomainToEntity(userDI)
	logger.Info("Saving user...",
		zap.String("UserRepositoryInterface", "Create"))

	_, err := u.sql.Exec("insert into usuario values (?,?,?,?,?)",
		userDI.GetID(),
		userDI.GetName(),
		userDI.GetEmail(),
		userDI.GetPassword(),
		userDI.GetAge(),
	)

	if err != nil {
		logger.Error("Erryr trying to save user",
			err, zap.String("UserRepositoryInterface", "Create"))
	}

	userDomain := entity.UserEntityToDomain(*userEntity)

	return userDomain, nil
}
