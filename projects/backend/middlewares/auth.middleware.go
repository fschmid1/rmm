package middlewares

import (
	"github.com/fes111/rmm/projects/backend/controller"
	"github.com/gofiber/fiber/v2"
)

func JwtAuth(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	if cookie == "" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	verify, user := controller.VerifyUserJWT(cookie)
	if verify {
		c.Locals("user", user)
		return c.Next()
	}

	return c.SendStatus(fiber.StatusUnauthorized)
}
