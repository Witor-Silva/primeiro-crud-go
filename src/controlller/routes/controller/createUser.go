package controller

import (
	"github.com/Witor-Silva/primeiro-crud-go/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	err := rest_err.NewBadRequestError("Você está chamando a rota de maneira errada!")
	c.JSON(err.Code, err)
}
