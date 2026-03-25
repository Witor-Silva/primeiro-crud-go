package request

import "github.com/Witor-Silva/primeiro-crud-go/src/model"

type UserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"Password" binding:"required,min=6"`
	Name     string `json:"name" binding:"required,min=3"`
	Age      int8   `json:"age" binding:"required,min=18,max=100"`
}

func (ur *UserRequest) ToDomain() model.UserDomainInterface {
	return model.NewUserDomain(
		ur.Email,
		ur.Password,
		ur.Name,
		ur.Age,
	)
}
