package model

import (
	"crypto/md5"

	"github.com/Witor-Silva/primeiro-crud-go/src/configuration/rest_err"
)

func NewUserDomain(
	email, Password, name string,
	age int8,
) *UserDomain {
	return &UserDomain{
		email, Password, name, age,
	}
}

type UserDomain struct {
	Email    string
	Password string
	Name     string
	Age      int8
}

func (ud *UserDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = string(hash.Sum(nil))
}

type UserDomainInterface interface {
	CreateUser() *rest_err.RestErr
	UpdateUser(string) *rest_err.RestErr
	FindUser(string) (*UserDomain, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
