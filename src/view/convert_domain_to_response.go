package view

import (
	"github.com/Witor-Silva/primeiro-crud-go/src/controlller/routes/controller/model/response"
	"github.com/Witor-Silva/primeiro-crud-go/src/model"
)

func ConvertDomainToResponse(userDomain model.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		ID:    userDomain.GetID(),
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
