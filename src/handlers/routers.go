package routers

import (
	"be/src/handlers/account"
	"be/src/handlers/auth"
	"be/src/handlers/group"
	"be/src/handlers/item"
	"github.com/gofiber/fiber/v2"
)

func Handlers(app *fiber.App) {
	api := app.Group("/api")
	auth.Router(api)
	account.Router(api)
	item.Router(api)
	group.Router(api)
}
