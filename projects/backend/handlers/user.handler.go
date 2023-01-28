package handlers

import (
	"github.com/fes111/rmm/libs/go/models"
	"github.com/fes111/rmm/projects/backend/controller"
	"github.com/gofiber/fiber/v2"
)

func HandleUserUpdate(c *fiber.Ctx) error {
	userId := c.Locals("user").(models.User).ID

	user, err := controller.GetUserById(userId)
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

func HandleGetAllUsers(c *fiber.Ctx) error {
	users, err := controller.GetAllUsers()
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	return c.Status(200).JSON(users)
}

func HandleGetProfile(c *fiber.Ctx) error {
	userId := c.Locals("user").(models.User).ID

	user, err := controller.GetUserById(userId)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}

	return c.Status(200).JSON(user)
}
