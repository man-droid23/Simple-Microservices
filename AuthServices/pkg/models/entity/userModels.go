package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email" form:"email" gorm:"unique;not null;type:varchar(100)" validate:"required,email"`
	Password string `json:"-" form:"password" validate:"required,min=6"  gorm:"not null;column:password"`
	Name     string `json:"name" form:"name" gorm:"not null;type:varchar(100)" validate:"required,min=3"`
	Role     string `json:"role" form:"role" gorm:"not null;type:varchar(100)" validate:"required"`
}
