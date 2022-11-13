package main

import (
	"festech.de/rmm/backend/config"
	"festech.de/rmm/backend/handlers"
	"festech.de/rmm/backend/models"
	"festech.de/rmm/backend/socket"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config.Connect()

	socket.RegisterWebsocketRoute(app)


	app.Post("/start/:id", func(c *fiber.Ctx) error {
		socket.SendMessage(c.Params("id"), models.SocketEvent{
			Event: "start-usage",
		})
		return c.SendString("sending")
	})

	app.Post("/stop/:id", func(c *fiber.Ctx) error {
		socket.SendMessage(c.Params("id"), models.SocketEvent{
			Event: "stop-usage",
		})
		return c.SendString("sending")
	})

	app.Get("/devices", handlers.GetDevices)
	app.Get("/devices/:id", handlers.GetDevice)
	app.Post("/devices", handlers.AddDevice)
	app.Patch("/devices/:id", handlers.UpdateDevice)
	app.Delete("/devices/:id", handlers.RemoveDevice)

	app.Listen(":" + config.Getenv("PORT", "8080"))
}
