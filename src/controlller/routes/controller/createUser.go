package controller

import (
	"fmt"

	"github.com/Witor-Silva/primeiro-crud-go/src/configuration/rest_err"
	"github.com/Witor-Silva/primeiro-crud-go/src/controlller/routes/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var userResquest request.UserRequest

	if err := c.ShouldBindJSON(&userResquest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("There are some incorrect fields, error=%s", err.Error))
		c.JSON(restErr.Code, restErr)
		return
	}
	fmt.Println(userResquest)
}
