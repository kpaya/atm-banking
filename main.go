package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"

	"github.com/kpaya/atm-banking/src/infra"
	"github.com/kpaya/atm-banking/src/router"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Panicf("it's not possible to load the .env file: %v", err.Error())
	}
}

func main() {
	fmt.Println("Hello World - ATM Banking Project")

	// Initilize the database
	infra.InitializeDB()

	// Initialize the router
	router.Initialize()

}
