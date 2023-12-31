package customer

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type TransactionStatus string

const (
	SUCCESS     TransactionStatus = "Success"
	FAILURE     TransactionStatus = "Failure"
	INTERRUPTED TransactionStatus = "Interrupted"
	FULL        TransactionStatus = "Full"
	PARTIAL     TransactionStatus = "Partial"
	NONE        TransactionStatus = "None"
)

type ITransaction interface {
	DefineStatus(status TransactionStatus)
	CheckStatus() TransactionStatus
}

type Transaction struct {
	ID        string            `json:"id" validate:"required" gorm:"primaryKey,default:gen_random_uuid()"`
	Status    TransactionStatus `json:"status" validate:"required"`
	CreatedAt time.Time         `json:"created_at"`
}

func (t *Transaction) CheckStatus() TransactionStatus {
	return t.Status
}

func (t *Transaction) DefineStatus(status TransactionStatus) {
	t.Status = status
}

// Withdrawal

type Withdrawal struct {
	Transaction
	Amount float64 `json:"amount" validate:"required,number,gt=0"`
}

func NewWithdrawal(transaction Transaction, amount float64) *Withdrawal {
	withdrawal := new(Withdrawal)
	withdrawal.Transaction = transaction
	withdrawal.Amount = amount
	if err := validator.New().Struct(withdrawal); err != nil {
		return nil
	}
	return withdrawal
}

// Deposit

type IDeposit interface {
	CheckAmmount() float64
}

type Deposit struct {
	Transaction
	Amount float64 `json:"amount" validate:"required,number,gt=0"`
}

func (d *Deposit) CheckAmmount() float64 {
	return d.Amount
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
	if err := validator.New().Struct(checkDeposit); err != nil {
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
	if err := validator.New().Struct(cashDeposit); err != nil {
		return nil
	}
	return cashDeposit
}

// BalanceInquiry

type BalanceInquiry struct {
	Transaction
	AccountId string `json:"account_id" validate:"required"`
}

func NewBalanceInquiry(accountId string, transaction Transaction) *BalanceInquiry {
	balanceInquiry := new(BalanceInquiry)
	balanceInquiry.Transaction = transaction
	balanceInquiry.AccountId = accountId
	if err := validator.New().Struct(balanceInquiry); err != nil {
		return nil
	}
	return balanceInquiry
}

// Transfer

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
	if err := validator.New().Struct(transfer); err != nil {
		return nil
	}
	return transfer
}
