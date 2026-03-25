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

	if err := uc.service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
		return
	}

	// log o user criado
	logger.Info("User created successfully",
		zap.String("journey", "CreateUser"),
		zap.String("email", domain.GetEmail()),
		zap.String("name", domain.GetName()),
		zap.Int8("age", domain.GetAge()),
	)
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domain))
}
