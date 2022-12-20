package main

import (
	"festech.de/rmm/backend/config"
	"festech.de/rmm/backend/handlers"
	"festech.de/rmm/backend/socket"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func main() {
	app := fiber.New()

	config.Connect()

	socket.RegisterWebsocketRoute(app)

	app.Post("/auth/login", handlers.HandleLogin)
	app.Post("/auth/signup", handlers.HandleSignUp)

	deviceRouter := app.Group("/devices")
	deviceRouter.Use(jwtware.New(config.JWT_CONFIG))

	deviceRouter.Post("/functions", socket.FunctionsHandler)

	deviceRouter.Post("/token", handlers.AddDeviceToken)

	deviceRouter.Get("/", handlers.GetDevices)
	deviceRouter.Get("/:id", handlers.GetDevice)
	deviceRouter.Post("/", handlers.AddDevice)
	deviceRouter.Patch("/", handlers.UpdateDevice)
	deviceRouter.Delete("/:id", handlers.RemoveDevice)

	app.Listen(":" + config.Getenv("PORT", "8080"))
}
