package handlers

import (
	"github.com/fes111/rmm/projects/backend/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func HandleUserUpdate(c *fiber.Ctx) error {
	userId := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["user"].(map[string]interface{})["id"]

	user, err := controller.GetUserById(uint(userId.(float64)))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	err = controller.UpdateUser(user)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	return c.Status(200).JSON(user)
}

func HandleGetProfile(c *fiber.Ctx) error {
	userId := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)["user"].(map[string]interface{})["id"]

	user, err := controller.GetUserById(uint(userId.(float64)))
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}

	return c.Status(200).JSON(user)
}
