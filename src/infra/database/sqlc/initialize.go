package db

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/jackc/pgx/v5"
)

var once sync.Once
var singleInstance *single

type single struct {
	Db *pgx.Conn
}

func InitializeDB() *pgx.Conn {
	if singleInstance == nil {
		once.Do(func() {
			// Get the database name from the .env file
			var dbName = os.Getenv("DB_NAME") + ".db"

			// Open the database
			db, err := pgx.Connect(context.Background(), fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_TIMEZONE")))
			if err != nil {
				log.Panicf("it's not possible to open the database: %v", err.Error())
			}

			// Create database if not exists
			if dbFileOpened, err := os.Open(dbName); errors.Is(err, os.ErrNotExist) {
				// Close the file after the function ends
				defer dbFileOpened.Close()

				createdDbFile, err := os.Create(dbName)
				if err != nil {
					log.Panicf("it's not possible to create the database: %v", err.Error())
				}

				// Close the file after the function ends
				defer createdDbFile.Close()

				// Read the schema file
				schemaContent, err := os.ReadFile("src/infra/database/sqlc/build/schema.sql")
				if err != nil {
					log.Panicf("it's not possible to read the schema file: %v", err.Error())
				}

				_, err = db.Exec(context.Background(), string(schemaContent))
				if err != nil {
					log.Panicf("it couldn't initialize database: %v", err.Error())
				}

				singleInstance = &single{Db: db}
			}
		})
	}
	return singleInstance.Db
}
