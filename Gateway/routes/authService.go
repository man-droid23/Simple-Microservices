package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/simple-microservices-gateway/controllers"
)

func AuthService(app *fiber.App) {
	auth := app.Group("/api")
	auth.Post("/login", controllers.Login)
	auth.Post("/register", controllers.Register)
}
