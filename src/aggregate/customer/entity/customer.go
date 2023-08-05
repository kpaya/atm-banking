package customer

import (
	"github.com/go-playground/validator/v10"
	valueObject "github.com/kpaya/atm-banking/src/aggregate/value-object"
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
	Name            string               `json:"name"`
	Email           string               `json:"email"`
	Phone           string               `json:"phone"`
	Status          CustomerStatus       `json:"status"`
	Address         *valueObject.Address `json:"address"`
	Card            *Card
	CheckingAccount *CheckingAccount
	SavingsAccount  *SavingsAccount
}

func NewCustomer(name, email, phone string, address *valueObject.Address, card *Card, checkingAccount *CheckingAccount, savingsAccount *SavingsAccount) *Customer {
	customer := new(Customer)
	customer.Name = name
	customer.Email = email
	customer.Phone = phone
	customer.Status = ACTIVE
	customer.Address = address
	customer.Card = card
	customer.CheckingAccount = checkingAccount
	customer.SavingsAccount = savingsAccount
	if err := customer.Validate(); err != nil {
		return nil
	}
	return customer
}

func (c *Customer) Validate() error {
	return validator.New().Struct(c)
}
