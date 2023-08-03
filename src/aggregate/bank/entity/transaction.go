package entity

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type TransactionStatus string

type TransactionType interface {
	Transaction | Withdrawal | Deposit | CheckDeposit | CashDeposit | BalanceInquiry | Transfer
}

const (
	SUCCESS TransactionStatus = "Success"
	FAILURE TransactionStatus = "Failure"
	BLOCKED TransactionStatus = "Blocked"
	FULL    TransactionStatus = "Full"
	PARTIAL TransactionStatus = "Partial"
	NONE    TransactionStatus = "None"
)

func Validate[T TransactionType](transactionType *T) error {
	return validator.New().Struct(transactionType)
}

type Transaction struct {
	ID        string            `json:"id" validate:"required"`
	Status    TransactionStatus `json:"status" validate:"required"`
	CreatedAt time.Time         `json:"created_at"`
}

func NewTransaction(id string, status TransactionStatus, createdAt time.Time) *Transaction {
	transaction := new(Transaction)
	transaction.ID = id
	transaction.Status = status
	transaction.CreatedAt = createdAt
	if err := Validate[Transaction](transaction); err != nil {
		return nil
	}
	return transaction
}

type Withdrawal struct {
	Transaction
	Amount float64 `json:"amount" validate:"required,number,gt=0"`
}

func NewWithdrawal(transaction Transaction, amount float64) *Withdrawal {
	withdrawal := new(Withdrawal)
	withdrawal.Transaction = transaction
	withdrawal.Amount = amount
	if err := Validate[Withdrawal](withdrawal); err != nil {
		return nil
	}
	return withdrawal
}

type Deposit struct {
	Transaction
	Amount float64 `json:"amount" validate:"required,number,gt=0"`
}

func NewDeposit(transaction Transaction, amount float64) *Deposit {
	deposit := new(Deposit)
	deposit.Transaction = transaction
	deposit.Amount = amount
	if err := Validate[Deposit](deposit); err != nil {
		return nil
	}
	return deposit
}

type CheckDeposit struct {
	Deposit
	CheckNumber string `json:"check_number" validate:"required"`
	BankCode    string `json:"bank_code" validate:"required"`
}

func NewCheckDeposit(checkNumber, bankCode string, deposit Deposit) *CheckDeposit {
	checkDeposit := new(CheckDeposit)
	checkDeposit.Deposit = deposit
	checkDeposit.CheckNumber = checkNumber
	checkDeposit.BankCode = bankCode
	if err := Validate[CheckDeposit](checkDeposit); err != nil {
		return nil
	}
	return checkDeposit
}

type CashDeposit struct {
	Deposit
	CashDepositLimit float64 `json:"cash_deposit_limit" validate:"required,number,gte=0"`
}

func NewCashDeposit(cashDepositLimit float64, deposit Deposit) *CashDeposit {
	cashDeposit := new(CashDeposit)
	cashDeposit.Deposit = deposit
	cashDeposit.CashDepositLimit = cashDepositLimit
	if err := Validate[CashDeposit](cashDeposit); err != nil {
		return nil
	}
	return cashDeposit
}

type BalanceInquiry struct {
	Transaction
	AccountId string `json:"account_id" validate:"required"`
}

func NewBalanceInquiry(accountId string, transaction Transaction) *BalanceInquiry {
	balanceInquiry := new(BalanceInquiry)
	balanceInquiry.Transaction = transaction
	balanceInquiry.AccountId = accountId
	if err := Validate[BalanceInquiry](balanceInquiry); err != nil {
		return nil
	}
	return balanceInquiry
}

type Transfer struct {
	Transaction
	FromAccountNumber string  `json:"from_account_number" validate:"required"`
	ToAccountNumber   string  `json:"to_account_number" validate:"required"`
	Amount            float64 `json:"amount" validate:"required,number,gt=0"`
}

func NewTransfer(fromAccountNumber, toAccountNumber string, amount float64, transaction Transaction) *Transfer {
	transfer := new(Transfer)
	transfer.Transaction = transaction
	transfer.FromAccountNumber = fromAccountNumber
	transfer.ToAccountNumber = toAccountNumber
	transfer.Amount = amount
	if err := Validate[Transfer](transfer); err != nil {
		return nil
	}
	return transfer
}
