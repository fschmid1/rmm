package main

import (
	"github.com/fes111/rmm/libs/go/helpers"
	"github.com/fes111/rmm/projects/backend/config"
	"github.com/fes111/rmm/projects/backend/handlers"
	"github.com/fes111/rmm/projects/backend/socket"
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

	userRouter := app.Group("/user")
	userRouter.Use(jwtware.New(config.JWT_CONFIG))

	userRouter.Get("/", handlers.HandleGetProfile)
	userRouter.Patch("/notifications/toggle", handlers.HandleToggleDeviceNotfication)
	userRouter.Get("/notifications", handlers.HandleGetDeviceNotfications)
	userRouter.Patch("/", handlers.HandleUserUpdate)

	deviceRouter := app.Group("/devices")
	deviceRouter.Use(jwtware.New(config.JWT_CONFIG))

	deviceRouter.Post("/functions", socket.FunctionsHandler)

	deviceRouter.Post("/tokens", handlers.AddDeviceToken)
	deviceRouter.Get("/tokens", handlers.GetDeviceTokens)
	deviceRouter.Delete("/tokens/:id", handlers.DeleteDeviceToken)

	deviceRouter.Get("/", handlers.GetDevices)
	deviceRouter.Get("/:id", handlers.GetDevice)
	deviceRouter.Post("/", handlers.AddDevice)
	deviceRouter.Patch("/", handlers.UpdateDevice)
	deviceRouter.Delete("/:id", handlers.RemoveDevice)

	app.Listen(":" + helpers.Getenv("PORT", "8080"))
}
