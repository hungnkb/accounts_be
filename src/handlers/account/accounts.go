package account

import (
	"github.com/gofiber/fiber/v2"
)

func GetAll(c *fiber.Ctx) error {
	a := "hoho"
	return c.JSON(a)
}
