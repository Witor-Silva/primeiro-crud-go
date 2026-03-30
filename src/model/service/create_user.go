package service

import (
	"fmt"

	"github.com/Witor-Silva/primeiro-crud-go/src/configuration/logger"
	"github.com/Witor-Silva/primeiro-crud-go/src/configuration/rest_err"
	"github.com/Witor-Silva/primeiro-crud-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("Init createUser model", zap.String("Journey", "createUser"))

	userDomain.EncryptPassword()

	userDomainRepository, err := ud.userRepository.CreateUser(userDomain)
	if err != nil {
		return nil, err
	}
	fmt.Println(userDomain.GetPassword())
	return userDomainRepository, nil
}
