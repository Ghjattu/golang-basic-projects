package main

import (
	"crm-basic-fiber/model"
	"log"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	app.Get("/api/lead", model.GetLeads)
	app.Get("/api/lead/:id", model.GetLeadByID)
	app.Post("/api/lead", model.CreateLead)
	app.Delete("/api/lead/:id", model.DeleteLead)
}

func main() {
	app := fiber.New()
	RegisterRoutes(app)

	log.Fatalln(app.Listen(":8080"))
}
