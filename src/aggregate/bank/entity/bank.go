package entity

import "github.com/go-playground/validator/v10"

type Bank struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

func NewBank(name, code string) *Bank {
	bank := new(Bank)
	bank.Name = name
	bank.Code = code
	if err := bank.Validate(); err != nil {
		return nil
	}
	return bank
}

func (a *Bank) Validate() error {
	return validator.New().Struct(a)
}
