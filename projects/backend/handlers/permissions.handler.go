package handlers

import (
	"log"
	"strconv"

	"github.com/fes111/rmm/libs/go/models"
	"github.com/fes111/rmm/projects/backend/controller"
	"github.com/gofiber/fiber/v2"
)

func HandleGetDevicePermissions(c *fiber.Ctx) error {
	id, _ := strconv.ParseUint(c.Params("id"), 10, 32)
	permissions, err := controller.GetDevicePermissions(uint(id))
	if err != nil {
		log.Println(err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}
	return c.JSON(permissions)
}

func HandleUpdateDevicePermissions(c *fiber.Ctx) error {
	userId := c.Locals("user").(models.User).ID
	permissions := models.DevicePermissions{}
	err := c.BodyParser(&permissions)
	if err != nil {
		log.Println(err)
		return c.Status(400).JSON(fiber.Map{
			"error": "Bad Request",
		})
	}
	permission, err2 := controller.GetDevicePermissionsByUserId(permissions.DeviceID, uint64(userId))
	if err2 != nil || !permission.ChangePermissions {
		return c.SendStatus(403)
	}
	permissions, err = controller.UpdateDevicePermissions(permissions)
	if err != nil {
		log.Println(err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}
	return c.JSON(permissions)
}

func HandleDeleteDevicePermissions(c *fiber.Ctx) error {
	id, _ := strconv.ParseUint(c.Params("id"), 10, 32)
	err := controller.DeleteDevicePermissions(uint(id))
	if err != nil {
		log.Println(err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}
	return c.SendStatus(200)
}
