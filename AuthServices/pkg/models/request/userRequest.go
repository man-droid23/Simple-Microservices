package request

import "github.com/go-playground/validator/v10"

type UserRequest struct {
	Email    string `json:"email" form:"email" gorm:"unique;not null;type:varchar(100)" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=6"  gorm:"not null;column:password"`
	Name     string `json:"name" form:"name" gorm:"not null;type:varchar(100)" validate:"required,min=3"`
	Role     string `json:"role" form:"role" gorm:"not null;type:varchar(100)" validate:"required"`
}

func (u *UserRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
