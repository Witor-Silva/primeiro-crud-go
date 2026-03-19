package controller

import (
	"net/http"

	"github.com/Witor-Silva/primeiro-crud-go/src/configuration/logger"
	"github.com/Witor-Silva/primeiro-crud-go/src/configuration/validation"
	"github.com/Witor-Silva/primeiro-crud-go/src/controlller/routes/controller/model/request"
	"github.com/Witor-Silva/primeiro-crud-go/src/controlller/routes/controller/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
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

	response := response.UserResponse{
		ID:    "Test",
		Name:  userRequest.Name,
		Email: userRequest.Email,
		Age:   userRequest.Age,
	}
	logger.Info("User created successfully",
		zap.String("journey", "CreateUser"),
	)

	c.JSON(http.StatusOK, response)
}
