package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/man-droid23/simple-microservices/pkg/auth"
)

func UserRoutes(app *fiber.App) {
	app.Get("/user", auth.GetAllUser)
	app.Get("/user/:id", auth.GetUser)
	app.Post("/user", auth.CreateUser)
	app.Put("/user/:id", auth.UpdateUser)
	app.Delete("/user/:id", auth.DeleteUser)
}
