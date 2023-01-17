package controllers

import (
	"bytes"
	"github.com/gofiber/fiber/v2"
	"github.com/simple-microservices-gateway/models/request"
	"io"
	"net/http"
)

func Login(c *fiber.Ctx) error {
	reqUser := new(request.Login)
	if err := c.BodyParser(reqUser); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"Message": "Cannot Parse Request",
			"Error":   err.Error(),
		})
	}
	errValidare := reqUser.Validate()
	if errValidare != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"Message": "Bad Request",
		})
	}
	body := []byte(`{"email": "` + reqUser.Email + `", "password": "` + reqUser.Password + `"}`)
	// Get token from auth service
	res, errRes := http.Post("http://localhost:3000/login", "application/json", bytes.NewBuffer(body))
	if errRes != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"Message": "Cannot Get Token",
			"Error":   errRes.Error(),
		})
	}
	defer res.Body.Close()
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"Message": "Cannot Read Response",
			"Error":   err.Error(),
		})
	}
	return c.Status(res.StatusCode).JSON(fiber.Map{
		"Message": "Login Success",
		"Token":   string(resBody),
	})
}

func Register(c *fiber.Ctx) error {
	return c.SendString("Register")
}
