package response

import "github.com/man-droid23/simple-microservices/pkg/models/entity"

type UserResponse struct {
	Email string `json:"email" form:"email" gorm:"unique;not null;type:varchar(100)" validate:"required,email"`
	Name  string `json:"name" form:"name" gorm:"not null;type:varchar(100)" validate:"required,min=3"`
	Role  string `json:"role" form:"role" gorm:"not null;type:varchar(100)" validate:"required"`
}

func NewUserResponse(user entity.User) UserResponse {
	return UserResponse{
		Email: user.Email,
		Name:  user.Name,
		Role:  user.Role,
	}
}

func NewUserListResponse(user []entity.User) []UserResponse {
	var userList []UserResponse
	for _, user := range user {
		userList = append(userList, NewUserResponse(user))
	}
	return userList
}
