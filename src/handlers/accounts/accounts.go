package accounts

import (
	"github.com/gofiber/fiber/v2"
)

func GetAll(c *fiber.Ctx) error {
	a := "hehe"
	return c.JSON(a)
}
