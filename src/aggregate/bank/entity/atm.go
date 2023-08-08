package entity

import (
	"github.com/go-playground/validator/v10"
	valueObject "github.com/kpaya/atm-banking/src/aggregate/value-object"
)

type ATM struct {
	ID        string `json:"id" validate:"required" gorm:"primaryKey,default:gen_random_uuid()"`
	AddressID string
	Address   *valueObject.Address `json:"address" validate:"required" gorm:"foreignKey:AddressID"`
}

func NewATM(id string, address *valueObject.Address) *ATM {
	atm := new(ATM)
	atm.ID = id
	atm.Address = address
	if err := atm.Validate(); err != nil {
		return nil
	}
	return atm
}

func (a *ATM) Validate() error {
	return validator.New().Struct(a)
}
