package customer

import (
	"github.com/go-playground/validator/v10"
)

type IAccount interface {
	AddMoney(amount float64) error
	WithdrawMoney(amount float64) error
	GetAccountNumber() string
	GetTotalBalance() float64
	GetAvaliableBalance() float64
}

// Account Area
type Account struct {
	ID               string  `gorm:"primaryKey,default:gen_random_uuid()" json:"id"`
	AccountNumber    string  `json:"account_number" validate:"required"`
	TotalBalance     float64 `json:"total_balance" validate:"required,number"`
	AvaliableBalance float64 `json:"avaliable_balance" validate:"required,number"`
}

func (account *Account) AddMoney(amount float64) error {
	account.TotalBalance += amount
	return nil
}

func (account *Account) WithdrawMoney(amount float64) error {
	account.TotalBalance -= amount
	return nil
}

func (account *Account) GetAccountNumber() string {
	return account.AccountNumber
}

func (account *Account) GetTotalBalance() float64 {
	return account.TotalBalance
}

func (account *Account) GetAvaliableBalance() float64 {
	return account.AvaliableBalance
}

// SavingsAccount Area
type SavingsAccount struct {
	Account
	WithdrawLimit float64 `json:"withdraw_limit"`
}

func NewSavingsAccount(accountNumber string, totalBalance, avaliableBalance, withdrawLimit float64) IAccount {
	savingsAccount := new(SavingsAccount)
	savingsAccount.AccountNumber = accountNumber
	savingsAccount.TotalBalance = totalBalance
	savingsAccount.AvaliableBalance = avaliableBalance
	savingsAccount.WithdrawLimit = withdrawLimit
	if err := validator.New().Struct(savingsAccount); err != nil {
		return nil
	}
	return savingsAccount
}

// CheckingAccount Area

type CheckingAccount struct {
	Account
	DebitCardNumber string  `json:"debit_card_number"`
	WithdrawLimit   float64 `json:"withdraw_limit"`
}

func NewCheckingAccount(accountNumber string, totalBalance, avaliableBalance float64, debitCardNumber string) IAccount {
	checkingAccount := new(CheckingAccount)
	checkingAccount.AccountNumber = accountNumber
	checkingAccount.TotalBalance = totalBalance
	checkingAccount.AvaliableBalance = avaliableBalance
	checkingAccount.DebitCardNumber = debitCardNumber
	if err := validator.New().Struct(checkingAccount); err != nil {
		return nil
	}
	return checkingAccount
}
