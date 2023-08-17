package main

import (
	"log"

	"github.com/joho/godotenv"

	infra "github.com/kpaya/atm-banking/src/infra/database"
	"github.com/kpaya/atm-banking/src/router"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Panicf("it's not possible to load the .env file: %v", err.Error())
	}
}

func main() {
	// Initilize the database
	infra.NewDatabaseInstance()

	// Initialize the router
	router.Initialize()

}
