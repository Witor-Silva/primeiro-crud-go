package service

import (
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
		logger.Error(
			"Error trying to create user", err, zap.String("Journey", "CreateUser"))
		return nil, err
	}
	logger.Info(
		"User service executed successfully",
		zap.String("userId", userDomainRepository.GetID()),
		zap.String("Journey", "CreateUser"))
	return userDomainRepository, nil
}
