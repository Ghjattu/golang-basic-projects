package main

import (
	"jwt-authorization/controllers"
	"jwt-authorization/middlewares"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	app.Use("/api/users", middlewares.JWTAuthorization())
	app.Get("/api/users", controllers.GetAllUsers)

	app.Listen(":8080")
}
