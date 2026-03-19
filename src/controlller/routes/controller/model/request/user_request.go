package request

type UserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"Password" binding:"required,min=6,containsAny=!@#$%^&*"`
	Name     string `json:"name" binding:"required,min=3"`
	Age      int8   `json:"age" binding:"required,min=18,max=100"`
}
