package routers

import (
	"be/src/handlers/account"
	"be/src/handlers/auth"

	"github.com/gofiber/fiber/v2"
)

func Handlers(app *fiber.App) {
	api := app.Group("/api")
	auth.Router(api)
	account.Router(api)
}
