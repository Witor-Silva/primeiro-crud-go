package controller

import (
	"github.com/Witor-Silva/primeiro-crud-go/src/model/service"
	"github.com/gin-gonic/gin"
)

func NewUserControllerInterface(serviceInterface service.UserDomainService) *userControllerInterface {
	return &userControllerInterface{
		service: serviceInterface,
	}
}

type UserControllerInterface interface {
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	FindUserByEmail(c *gin.Context)
	FindUserById(c *gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}
