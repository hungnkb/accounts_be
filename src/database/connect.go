package connection

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Postgres() *gorm.DB {
	dsn := "host=localhost user=postgres password=123456 dbname=test port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
