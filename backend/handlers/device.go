package handlers

import (
	"festech.de/rmm/backend/config"
	"festech.de/rmm/backend/models"
	"github.com/gofiber/fiber/v2"
)

func GetDevices(c *fiber.Ctx) error {
	var devices []models.Device

	config.Database.Find(&devices)
	return c.Status(200).JSON(devices)
}

func GetDevice(c *fiber.Ctx) error {
	id := c.Params("id")

	var device models.Device

	result := config.Database.Find(&device, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(&device)
}

func AddDevice(c *fiber.Ctx) error {
	device := new(models.Device)

	if err := c.BodyParser(device); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	if result := config.Database.Create(&device); result.Error != nil {
		return c.Status(400).SendString(result.Error.Error())
	}
	return c.Status(201).JSON(device)
}

func UpdateDevice(c *fiber.Ctx) error {
	device := new(models.Device)
	id := c.Params("id")

	if err := c.BodyParser(device); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	config.Database.Where("id = ?", id).Updates(&device)
	return c.Status(200).JSON(device)
}

func RemoveDevice(c *fiber.Ctx) error {
	id := c.Params("id")
	var device models.Device

	result := config.Database.Delete(&device, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}
