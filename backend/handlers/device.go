package handlers

import (
	"errors"

	"festech.de/rmm/backend/config"
	"festech.de/rmm/backend/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetDevices(c *fiber.Ctx) error {
	var devices []models.Device

	result := config.Database.Preload(clause.Associations).Find(&devices)

	if result.Error != nil {
		return c.SendStatus(500)
	}
	return c.Status(200).JSON(devices)
}

func GetDevice(c *fiber.Ctx) error {
	id := c.Params("id")
	mac := c.Query("mac")

	var device models.Device
	var result *gorm.DB
	if mac != "" {
		systemInfo := GetSystemInfoByMacAddress(id)
		result = config.Database.Preload(clause.Associations).Where("system_info_id = ?", systemInfo.ID).First(&device)
	} else {
		result = config.Database.Preload(clause.Associations).Find(&device, id)
	}
	if result.Error != nil {
		return c.SendStatus(500)
	}
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

	if GetSystemInfoByMacAddress(device.SystemInfo.MacAddress).ID > 0 {
		return c.Status(400).SendString("Device with this mac address is already registered")
	}

	err := config.Database.Transaction(func(tx *gorm.DB) error {
		if result := tx.Create(&device); result.Error != nil {
			return result.Error
		}
		return nil
	})
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	return c.Status(200).JSON(device)
}

func UpdateDevice(c *fiber.Ctx) error {
	device := new(models.Device)

	if err := c.BodyParser(device); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	err := config.Database.Transaction(func(tx *gorm.DB) error {
		if result := tx.Updates(&device); result.Error != nil {
			return result.Error
		}
		if result := tx.Updates(&device.SystemInfo); result.Error != nil {
			return result.Error
		}
		return nil
	})
	if err != nil {
		return c.SendStatus(400)
	}
	return c.Status(200).JSON(device)
}

func RemoveDevice(c *fiber.Ctx) error {
	id := c.Params("id")
	var device models.Device

	result := config.Database.Delete(&device, id)

	if result.Error != nil {
		return c.SendStatus(500)
	}

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}

func GetDeviceByDeviceId(id string) (models.Device, error) {
	var device models.Device

	result := config.Database.Where("device_id = ?", id).First(&device)

	if result.RowsAffected == 0 {
		return models.Device{}, errors.New("not found")
	}

	return device, nil
}

func SetDeviceConnected(id string, connected bool) (bool, error) {
	device, err := GetDeviceByDeviceId(id)
	if err != nil {
		return false, errors.New("not found")
	}

	if result := config.Database.Model(&device).Update("connected", connected); result.Error != nil {
		return false, errors.New("db error")
	}
	return true, nil
}
