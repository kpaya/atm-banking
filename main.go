package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kpaya/atm-banking/src/router"

	db "github.com/kpaya/atm-banking/src/infra/database/sqlc"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Panicf("it's not possible to load the .env file: %v", err.Error())
	}
}

func main() {

	// Initialize the database
	db.InitializeDB()

	// Initialize the router
	router.Initialize()

}
