package handlers

import (
	"log"

	"github.com/fes111/rmm/libs/go/models"
	"github.com/fes111/rmm/projects/backend/config"
	"github.com/fes111/rmm/projects/backend/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func AddDeviceToken(c *fiber.Ctx) error {
	deviceToken := new(models.DeviceToken)
	var err error

	if err := c.BodyParser(deviceToken); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["user"].(map[string]interface{})["id"]
	deviceToken.UserID = uint64(id.(float64))
	deviceToken.Token, err = controller.GenerateDeviceJWT(*deviceToken)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	err = config.Database.Transaction(func(tx *gorm.DB) error {
		if result := tx.Create(&deviceToken); result.Error != nil {
			return result.Error
		}
		return nil
	})
	if err != nil {
		log.Println(err.Error())
		return c.Status(400).SendString(err.Error())
	}

	return c.Status(200).JSON(deviceToken)
}

func GetDeviceTokens(c *fiber.Ctx) error {
	var tokens []models.DeviceToken
	userId := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["user"].(map[string]interface{})["id"]
	result := config.Database.Preload(clause.Associations).Find(&tokens).Where("user_id = ?", userId)

	if result.Error != nil {
		return c.SendStatus(500)
	}
	return c.Status(200).JSON(tokens)
}

func DeleteDeviceToken(c *fiber.Ctx) error {
	id := c.Params("id")
	var token models.DeviceToken
	result := config.Database.Preload(clause.Associations).Find(&token, id)

	if result.Error != nil {
		return c.SendStatus(500)
	}
	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	result = config.Database.Delete(&token)
	if result.Error != nil {
		return c.SendStatus(500)
	}
	return c.SendStatus(200)
}

func GetDevices(c *fiber.Ctx) error {
	var devices []models.Device

	userId := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["user"].(map[string]interface{})["id"]
	result := config.Database.Preload(clause.Associations).Where("id IN (?)", config.Database.Table("device_permissions").Select("device_id").Where("user_id = ?", userId)).Find(&devices)

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
		systemInfo := controller.GetSystemInfoByMacAddress(id)
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

	if controller.GetSystemInfoByMacAddress(device.SystemInfo.MacAddress).ID > 0 {
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
