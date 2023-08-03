package customer

import "github.com/go-playground/validator/v10"

type AccountType interface {
	account | SavingsAccount | CheckingAccount
}

func Validate[T AccountType](accountType *T) error {
	return validator.New().Struct(accountType)
}

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
	if err := Validate[account](newAccount); err != nil {
		return nil
	}
	return newAccount
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
	if err := Validate[SavingsAccount](savingsAccount); err != nil {
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
	if err := Validate[CheckingAccount](checkingAccount); err != nil {
		return nil
	}
	return checkingAccount
}
