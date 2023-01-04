package handlers

import (
	"github.com/fes111/rmm/projects/backend/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func HandleToggleDeviceNotfication(c *fiber.Ctx) error {
	payload := struct {
		DeviceID uint `json:"deviceID"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return c.SendStatus(500)
	}
	userId := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["user"].(map[string]interface{})["id"].(float64)
	controller.ToggleDeviceNotification(uint(userId), payload.DeviceID)
	return c.JSON(map[string]bool{
		"success": true,
	})
}

func HandleGetDeviceNotfications(c *fiber.Ctx) error {
	userId := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["user"].(map[string]interface{})["id"].(float64)
	notifications := controller.GetDeviceNotificationsByUserID(uint(userId))
	return c.JSON(notifications)
}
