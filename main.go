package main

import (
	httpMessageEnum "be/src/common/httpEnum/httpMessage"
	connection "be/src/database"
	routers "be/src/handlers"
	"be/src/models/accountModel"
	credentialModel "be/src/models/model"
	"errors"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

type ErrorHandler struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading environment file")
	}
}

func main() {
	db := connection.Postgres()
	db.AutoMigrate(&accountModel.Account{}, &credentialModel.Credential{})
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}
			err = c.Status(code).JSON(&ErrorHandler{
				Status:  code,
				Message: e.Message,
			})
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(&ErrorHandler{
					Status:  fiber.StatusInternalServerError,
					Message: httpMessageEnum.WRONG,
				})
			}
			return nil
		},
	})
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))
	routers.Handlers(app)
	PORT := ":" + os.Getenv("PORT")
	log.Fatal(app.Listen(PORT))
}
