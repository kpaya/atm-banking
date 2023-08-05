package customer

import "github.com/go-playground/validator/v10"

type account struct {
	AccountNumber    string  `json:"account_number" validate:"required"`
	TotalBalance     float64 `json:"total_balance" validate:"required,number"`
	AvaliableBalance float64 `json:"avaliable_balance" validate:"required,number"`
}

func NewAccount(accountNumber string, totalBalance, avaliableBalance float64) *account {
	newAccount := new(account)
	newAccount.AccountNumber = accountNumber
	newAccount.TotalBalance = totalBalance
	newAccount.AvaliableBalance = avaliableBalance
	if err := validator.New().Struct(newAccount); err != nil {
		return nil
	}
	return newAccount
}

func (a *account) CheckBalance() float64 {
	return a.TotalBalance
}

type SavingsAccount struct {
	account
	WithdrawLimit float64 `json:"withdraw_limit"`
}

func NewSavingAccount(accountNumber string, totalBalance, avaliableBalance, withdrawLimit float64) *SavingsAccount {
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

type CheckingAccount struct {
	account
	DebitCardNumber string `json:"debit_card_number"`
}

func NewCheckingAccount(accountNumber string, totalBalance, avaliableBalance float64, debitCardNumber string) *CheckingAccount {
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
