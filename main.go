package main

import (
	httpMessageEnum "be/src/common/httpEnum/httpMessage"
	connection "be/src/database"
	routers "be/src/handlers"
	"be/src/models/accountModel"
	"be/src/models/categoryModel"
	groupModel "be/src/models/groupItemModel"
	"be/src/models/itemModel"
	credentialModel "be/src/models/model"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
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

func handlePanic(c *fiber.Ctx, err interface{}) {
	fmt.Println("Recovered from panic ===", err)
	c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": err,
	})
}

func main() {
	db := connection.Postgres()
	db.AutoMigrate(&accountModel.Account{}, &credentialModel.Credential{}, &itemModel.Item{}, &categoryModel.Category{}, &groupModel.Group{})
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
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	app.Use(cors.New())
	app.Use(func(c *fiber.Ctx) error {
		defer func() {
			if r := recover(); r != nil {
				handlePanic(c, r)
			}
		}()
		return c.Next()
	})

	loggerConfig := logger.Config{
		Format:     "[${time}] - ${ip}:${port} ${status} - ${method} ${path}\n",
		TimeFormat: "15:04:05",
		TimeZone:   "Local",
	}

	app.Use(logger.New(loggerConfig))

	routers.Handlers(app)
	PORT := ":" + os.Getenv("PORT")
	log.Fatal(app.Listen(PORT))
}
