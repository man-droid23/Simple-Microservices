package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/man-droid23/simple-microservices/pkg/db"
	"github.com/man-droid23/simple-microservices/pkg/models/entity"
	"github.com/man-droid23/simple-microservices/pkg/models/request"
	"github.com/man-droid23/simple-microservices/pkg/models/response"
	"github.com/man-droid23/simple-microservices/pkg/utils"
	"net/http"
)

func GetAllUser(ctx *fiber.Ctx) error {
	var user []entity.User
	result := db.DB.Find(&user).Error
	if result != nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
			"Message": "User Not Found",
			"Data":    nil,
		})
	}
	return ctx.Status(http.StatusFound).JSON(fiber.Map{
		"Message": "Data Found",
		"Data":    response.NewUserListResponse(user),
	})
}

func GetUser(ctx *fiber.Ctx) error {
	var user entity.User
	id := ctx.Params("id")
	result := db.DB.Find(&user, id).Error
	if result != nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
			"Message": "Data Not Found",
			"Error":   result.Error(),
		})
	}
	return ctx.Status(http.StatusFound).JSON(fiber.Map{
		"Message": "Data Found",
		"Data":    response.NewUserResponse(user),
	})
}

func CreateUser(ctx *fiber.Ctx) error {
	userReq := new(request.UserRequest)
	err := ctx.BodyParser(&userReq)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"Message": "Bad Request",
			"Error":   err.Error(),
		})
	}
	newUser := entity.User{
		Email:    userReq.Email,
		Name:     userReq.Name,
		Password: userReq.Password,
		Role:     userReq.Role,
	}
	hashPass, errHashPass := utils.HashPassword(userReq.Password)
	if errHashPass != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"Message": "Internal Server Error",
		})
	}
	newUser.Password = hashPass
	errCreate := db.DB.Create(&newUser).Error
	if errCreate != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"Message": "Create User Failed",
			"Error":   errCreate.Error(),
		})
	}
	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"Message": "Create User Success",
		"Data":    response.NewUserResponse(newUser),
	})
}

func UpdateUser(ctx *fiber.Ctx) error {
	var user entity.User
	id := ctx.Params("id")
	userReq := new(request.UserRequest)
	err := ctx.BodyParser(&userReq)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"Message": "Bad Request",
			"Error":   err.Error(),
		})
	}
	errValidate := userReq.Validate()
	if errValidate != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"Message": "Bad Request",
			"Error":   errValidate.Error(),
		})
	}
	result := db.DB.Find(&user, id).Error
	if result != nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
			"Message": "Data Not Found",
			"Error":   result.Error(),
		})
	}
	hashPass, errHashPass := utils.HashPassword(userReq.Password)
	if errHashPass != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"Message": "Internal Server Error",
		})
	}
	user.Email = userReq.Email
	user.Name = userReq.Name
	user.Password = hashPass
	user.Role = userReq.Role
	errUpdate := db.DB.Save(&user).Error
	if errUpdate != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"Message": "Update User Failed",
			"Error":   errUpdate.Error(),
		})
	}
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"Message": "Update User Success",
	})
}

func DeleteUser(ctx *fiber.Ctx) error {
	var user entity.User
	id := ctx.Params("id")
	result := db.DB.Find(&user, id).Error
	if result != nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
			"Message": "Data Not Found",
			"Error":   result.Error(),
		})
	}
	errDelete := db.DB.Delete(&user).Error
	if errDelete != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"Message": "Delete User Failed",
			"Error":   errDelete.Error(),
		})
	}
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"Message": "Delete User Success",
	})
}
