package dtos

type CreateUserDTO struct {
	Username string `json:"username" validate:"required,min=3,max=20"`
	Email    string `json:"email" validate:"required,email"`
}

type UpdateUserDTO struct {
	Username string `json:"username" validate:"omitempty,min=3,max=20"`
	Email    string `json:"email" validate:"omitempty,email"`
}
