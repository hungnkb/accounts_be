package main

import (
	connection "be/src/database"
	routers "be/src/handlers"
	"be/src/models/accountModel"
	credentialModel "be/src/models/model"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading environment file")
	}
}

func main() {
	fmt.Println("123123")
	db := connection.Postgres()
	db.AutoMigrate(&accountModel.Account{}, &credentialModel.Credential{})
	app := fiber.New()
	routers.Handlers(app)
	PORT := ":" + os.Getenv("PORT")
	log.Fatal(app.Listen(PORT))
}
