package connection

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Postgres() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading environment file")
	}
	dsn := os.Getenv("POSTGRES_STRING_CONNECTION")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
