package controller

import (
	"fmt"
	"log"

	"github.com/Witor-Silva/primeiro-crud-go/src/configuration/validation"
	"github.com/Witor-Silva/primeiro-crud-go/src/controlller/routes/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Error trying to unmarshal object, error=%s\n", err.Error())
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	fmt.Println(userRequest)
}
