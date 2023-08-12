package customer

import (
	"time"

	"github.com/go-playground/validator/v10"
	valueObject "github.com/kpaya/atm-banking/src/aggregate/value-object"
	"gorm.io/gorm"
)

type CustomerStatus string

const (
	ACTIVE      CustomerStatus = "Active"
	BLOCKED     CustomerStatus = "Blocked"
	BANNED      CustomerStatus = "Banned"
	COMPROMISED CustomerStatus = "Compromised"
	ARCHIVED    CustomerStatus = "Archived"
	CLOSED      CustomerStatus = "Closed"
	UNKNOWN     CustomerStatus = "Unknown"
)

type Customer struct {
	ID        string               `gorm:"primaryKey,default:gen_random_uuid()" json:"id"`
	Name      string               `json:"name"`
	Email     string               `json:"email"`
	Phone     string               `json:"phone"`
	Status    CustomerStatus       `json:"status"`
	AddressID string               `json:"address_id"`
	Address   *valueObject.Address `json:"address"`
	AccountID string               `json:"account_id" `
	Account   AccountType          `json:"account" gorm:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func NewCustomer(name, email, phone string, address *valueObject.Address, card *Card, account AccountType) *Customer {
	customer := new(Customer)
	customer.Name = name
	customer.Email = email
	customer.Phone = phone
	customer.Status = ACTIVE
	customer.Address = address
	customer.Account = account
	if err := customer.Validate(); err != nil {
		return nil
	}
	return customer
}

func (c *Customer) Validate() error {
	return validator.New().Struct(c)
}
