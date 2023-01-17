package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/simple-microservices-gateway/routes"
)

func main() {
	app := fiber.New()
	routes.AuthService(app)

	err := app.Listen(":8000")
	if err != nil {
		panic(err)
	}
}
