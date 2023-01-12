package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/man-droid23/simple-microservices/pkg/auth"
)

func AuthRoutes(app *fiber.App) {
	app.Post("/login", auth.LoginController)
	app.Post("/register", auth.RegisterController)
}
