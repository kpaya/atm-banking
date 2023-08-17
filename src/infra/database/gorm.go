package infra

import (
	"fmt"
	"log"
	"os"
	"sync"

	bankAgg "github.com/kpaya/atm-banking/src/aggregate/bank/entity"
	customerAgg "github.com/kpaya/atm-banking/src/aggregate/customer/entity"
	value_object "github.com/kpaya/atm-banking/src/aggregate/value-object"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var syncInstance sync.Once
var sigletonInstance *singletonDatabase

type singletonDatabase struct {
	Db *gorm.DB
}

func NewDatabaseInstance() *gorm.DB {

	if sigletonInstance == nil {
		syncInstance.Do(func() {
			db, err := gorm.Open(postgres.Open(
				fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_TIMEZONE"))),
				&gorm.Config{})

			if err != nil {
				log.Panicf("it's not possible to connect to the database: %v", err.Error())
			}

			// Migrate the schema
			err = db.AutoMigrate(&value_object.Address{}, bankAgg.ATM{}, &customerAgg.Card{}, &customerAgg.Account{}, &customerAgg.Customer{})
			if err != nil {
				log.Panicf("it's not possible to migrate database: %v", err.Error())
			}

			sigletonInstance = &singletonDatabase{Db: db}
			fmt.Println("db instance created")
		})
	}

	return sigletonInstance.Db
}
