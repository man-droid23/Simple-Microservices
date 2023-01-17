package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/man-droid23/simple-microservices/pkg/db"
	"github.com/man-droid23/simple-microservices/pkg/models/entity"
	"github.com/man-droid23/simple-microservices/pkg/models/request"
	"github.com/man-droid23/simple-microservices/pkg/utils"
	"net/http"
	"time"
)

func LoginController(c *fiber.Ctx) error {
	var user entity.User
	reqUser := new(request.LoginRequest)
	err := c.BodyParser(reqUser)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"Message": "Cannot Parse Request",
			"Error":   err.Error(),
		})
	}
	errValidate := reqUser.Validate()
	if errValidate != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"Message": "Bad Request",
		})
	}
	errLogin := db.DB.Where("email = ?", reqUser.Email).First(&user).Error
	if errLogin != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"Message": "Unauthorized",
		})
	}
	pass := utils.CheckPasswordHash(reqUser.Password, user.Password)
	if !pass {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"Message": "Invalid Password",
		})
	}
	claims := jwt.MapClaims{}
	claims["email"] = user.Email
	claims["name"] = user.Name
	claims["role"] = user.Role
	claims["exp"] = time.Now().Add(time.Minute * 2).Unix()
	token, errToken := GenerateJJWTToken(&claims)
	if errToken != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"Message": "Cannot Generate Token",
			"Error":   errToken.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(token)
}

func RegisterController(ctx *fiber.Ctx) error {
	reqUser := new(request.UserRequest)
	err := ctx.BodyParser(reqUser)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"Message": "Bad Request",
			"Error":   err.Error(),
		})
	}
	newUser := entity.User{
		Email:    reqUser.Email,
		Name:     reqUser.Name,
		Password: reqUser.Password,
		Role:     reqUser.Role,
	}
	hash, errHash := utils.HashPassword(newUser.Password)
	if errHash != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"Message": "Cannot Hash Password",
			"Error":   err.Error(),
		})
	}
	newUser.Password = hash
	errCreate := db.DB.Create(&newUser).Error
	if errCreate != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"Message": "Cannot Create User",
			"Error":   errCreate.Error(),
		})
	}
	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"Message": "Register Success",
	})
}

func RefreshController(ctx *fiber.Ctx) error {
	tokenValidation := ctx.Locals("token").(*jwt.Token)
	claims := tokenValidation.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Minute * 2).Unix()
	newToken, errToken := GenerateJJWTToken(&claims)
	if errToken != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"Message": "Cannot Generate Token",
			"Error":   errToken.Error(),
		})
	}
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"Message": "Refresh Token Success",
		"Token":   newToken,
	})
}

func LogoutController(ctx *fiber.Ctx) error {
	// Logout Logic
	makeTokenInvalid := ctx.Locals("token").(*jwt.Token)
	claims := makeTokenInvalid.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Minute * -2).Unix()
	_, errToken := GenerateJJWTToken(&claims)
	if errToken != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"Message": "Cannot Generate Token",
			"Error":   errToken.Error(),
		})
	}
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"Message": "Logout Success",
	})
}
