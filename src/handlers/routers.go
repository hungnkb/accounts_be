package routers

import (
	// "be/src/handlers/accounts"

	"be/src/handlers/accounts"

	"github.com/gofiber/fiber/v2"
)

func Handlers(app *fiber.App) {
	api := app.Group("/api")
	
	api.Get("/accounts", accounts.GetAll)
}
