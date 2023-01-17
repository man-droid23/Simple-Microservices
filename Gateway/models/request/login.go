package request

import "github.com/go-playground/validator/v10"

type Login struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=6"`
}

func (l *Login) Validate() error {
	validate := validator.New()
	return validate.Struct(l)
}
