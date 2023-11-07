package main

import (
	routers "be/src/handlers"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	routers.Handlers(app)
	log.Fatal(app.Listen(":5050"))
}
