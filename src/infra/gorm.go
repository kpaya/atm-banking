package infra

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open(
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_TIMEZONE"))),
		&gorm.Config{})

	if err != nil {
		log.Panicf("it's not possible to connect to the database: %v", err.Error())
	}

	// Migrate the schema
	// err = db.AutoMigrate()
	// if err != nil {
	// 	log.Panicf("it's not possible to migrate database: %v", err.Error())
	// }

	return db
}
