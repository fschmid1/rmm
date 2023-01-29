package main

import (
	"github.com/fes111/rmm/libs/go/helpers"
	"github.com/fes111/rmm/projects/backend/config"
	"github.com/fes111/rmm/projects/backend/handlers"
	"github.com/fes111/rmm/projects/backend/middlewares"
	"github.com/fes111/rmm/projects/backend/socket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	config.Connect()

	socket.RegisterWebsocketRoute(app)

	app.Post("/auth/login", handlers.HandleLogin)
	app.Post("/auth/signup", handlers.HandleSignUp)
	app.Post("/auth/refresh", handlers.HandleRefreshToken)
	app.Post("/auth/logout", handlers.HandleLogout)

	userRouter := app.Group("/user")
	userRouter.Use(middlewares.JwtAuth)

	userRouter.Get("/", handlers.HandleGetProfile)
	userRouter.Get("/all", handlers.HandleGetAllUsers)
	userRouter.Patch("/notifications/toggle", handlers.HandleToggleDeviceNotfication)
	userRouter.Get("/notifications", handlers.HandleGetDeviceNotfications)
	userRouter.Patch("/", handlers.HandleUserUpdate)

	deviceRouter := app.Group("/devices")
	deviceRouter.Use(middlewares.JwtAuth)

	deviceRouter.Post("/functions", socket.FunctionsHandler)

	deviceRouter.Post("/tokens", handlers.AddDeviceToken)
	deviceRouter.Get("/tokens", handlers.GetDeviceTokens)
	deviceRouter.Delete("/tokens/:id", handlers.DeleteDeviceToken)

	deviceRouter.Get("/:id/permissions", handlers.HandleGetDevicePermissions)
	deviceRouter.Patch("/:id/permissions", handlers.HandleUpdateDevicePermissions)
	deviceRouter.Delete("/:id/permissions", handlers.HandleDeleteDevicePermissions)

	deviceRouter.Get("/", handlers.GetDevices)
	deviceRouter.Get("/:id", handlers.GetDevice)
	deviceRouter.Post("/", handlers.AddDevice)
	deviceRouter.Patch("/", handlers.UpdateDevice)
	deviceRouter.Delete("/:id", handlers.RemoveDevice)

	app.Listen(":" + helpers.Getenv("PORT", "8080"))
}
