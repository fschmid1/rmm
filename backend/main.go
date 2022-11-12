package main

import (
	"festech.de/rmm/backend/config"
	"festech.de/rmm/backend/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config.Connect()

	app.Get("/devices", handlers.GetDevices)
	app.Get("/devices/:id", handlers.GetDevice)
	app.Post("/devices", handlers.AddDevice)
	app.Put("/devices/:id", handlers.UpdateDevice)
	app.Delete("/devices/:id", handlers.RemoveDevice)

	app.Listen(":" + config.Getenv("PORT", "8080"))
}
