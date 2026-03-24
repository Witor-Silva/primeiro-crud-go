package service

import (
	"github.com/Witor-Silva/primeiro-crud-go/src/configuration/rest_err"
	"github.com/Witor-Silva/primeiro-crud-go/src/model"
)

func (ud *userDomainService) FindUser(string) (
	model.UserDomainInterface, *rest_err.RestErr) {
	return nil, nil
}
func (ud *userDomainService) FindUserByID(id string) (
	model.UserDomainInterface, *rest_err.RestErr) {
	return nil, nil
}
