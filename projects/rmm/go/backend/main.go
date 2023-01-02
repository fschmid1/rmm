package main

import (
	"github.com/fes111/rmm/projects/rmm/go/backend/config"
	"github.com/fes111/rmm/projects/rmm/go/backend/handlers"
	"github.com/fes111/rmm/projects/rmm/go/backend/socket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	jwtware "github.com/gofiber/jwt/v3"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	config.Connect()

	socket.RegisterWebsocketRoute(app)

	app.Post("/auth/login", handlers.HandleLogin)
	app.Post("/auth/signup", handlers.HandleSignUp)

	deviceRouter := app.Group("/devices")
	deviceRouter.Use(jwtware.New(config.JWT_CONFIG))

	deviceRouter.Post("/functions", socket.FunctionsHandler)

	deviceRouter.Post("/tokens", handlers.AddDeviceToken)
	deviceRouter.Get("/tokens", handlers.AddDeviceToken)

	deviceRouter.Get("/", handlers.GetDevices)
	deviceRouter.Get("/:id", handlers.GetDevice)
	deviceRouter.Post("/", handlers.AddDevice)
	deviceRouter.Patch("/", handlers.UpdateDevice)
	deviceRouter.Delete("/:id", handlers.RemoveDevice)

	app.Listen(":" + config.Getenv("PORT", "8080"))
}
