package controller

import (
	"net/http"

	"github.com/Witor-Silva/primeiro-crud-go/src/configuration/logger"
	"github.com/Witor-Silva/primeiro-crud-go/src/configuration/validation"
	"github.com/Witor-Silva/primeiro-crud-go/src/controlller/routes/controller/model/request"
	"github.com/Witor-Silva/primeiro-crud-go/src/model"
	"github.com/Witor-Silva/primeiro-crud-go/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zap.String("journey", "CreateUser"),
	)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to unmarshal object", err,
			zap.String("journey", "CreateUser"),
		)
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := userRequest.ToDomain()

	domainResult, err := uc.service.CreateUser(domain)
	if err != nil {
		logger.Error(
			"Error trying to create user", err,
			zap.String("Journey", "CreateUser"))
		c.JSON(err.Code, err)
		return
	}

	// log o user criado
	logger.Info("User controller executed successfully",
		zap.String("userId", domainResult.GetID()),
	)
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}
