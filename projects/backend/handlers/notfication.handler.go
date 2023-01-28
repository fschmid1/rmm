package handlers

import (
	"github.com/fes111/rmm/libs/go/models"
	"github.com/fes111/rmm/projects/backend/controller"
	"github.com/gofiber/fiber/v2"
)

func HandleToggleDeviceNotfication(c *fiber.Ctx) error {
	payload := struct {
		DeviceID uint `json:"deviceID"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return c.SendStatus(500)
	}
	userId := c.Locals("user").(models.User).ID
	controller.ToggleDeviceNotification(userId, payload.DeviceID)
	return c.JSON(map[string]bool{
		"success": true,
	})
}

func HandleGetDeviceNotfications(c *fiber.Ctx) error {
	userId := c.Locals("user").(models.User).ID
	notifications := controller.GetDeviceNotificationsByUserID(uint(userId))
	return c.JSON(notifications)
}
