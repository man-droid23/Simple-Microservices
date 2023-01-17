package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/man-droid23/simple-microservices/api"
	"github.com/man-droid23/simple-microservices/pkg/db"
)

func init() {
	db.ConnectDB()
	db.Migration()
}

func main() {
	defer db.CloseDB()
	r := fiber.New()
	api.AuthRoutes(r)
	api.UserRoutes(r)
	err := r.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
