package dtos

type CreateUserDTO struct {
	Username string `json:"username" validate:"required,min=3,max=20"`
	Email    string `json:"password" validate:"required,email"`
}
