package middlewares

import (
	"github.com/fes111/rmm/projects/backend/controller"
	"github.com/gofiber/fiber/v2"
)

func JwtAuth(c *fiber.Ctx) error {
	header := c.Get("Authorization", "")
	if header == "" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	token := header[7:]
	verify, user := controller.VerifyUserJWT(token)
	if verify {
		c.Locals("user", user)
		return c.Next()
	}

	return c.SendStatus(fiber.StatusUnauthorized)
}
