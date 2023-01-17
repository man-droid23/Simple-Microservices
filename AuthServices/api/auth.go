package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/man-droid23/simple-microservices/pkg/auth"
	"github.com/man-droid23/simple-microservices/pkg/middleware"
)

func AuthRoutes(app *fiber.App) {
	app.Post("/login", auth.LoginController)
	app.Post("/register", auth.RegisterController)
	app.Get("/refresh", middleware.AuthMiddleware, auth.RefreshController)
	app.Get("/logout", middleware.AuthMiddleware, auth.LogoutController)
}
