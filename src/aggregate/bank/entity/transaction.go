package entity

import "time"

type TransactionStatus string

const (
	SUCCESS TransactionStatus = "Success"
	FAILURE TransactionStatus = "Failure"
	BLOCKED TransactionStatus = "Blocked"
	FULL    TransactionStatus = "Full"
	PARTIAL TransactionStatus = "Partial"
	NONE    TransactionStatus = "None"
)

type Transaction struct {
	ID        string            `json:"id"`
	Status    TransactionStatus `json:"status"`
	CreatedAt time.Time         `json:"created_at"`
}

func NewTransaction(id string, status TransactionStatus, createdAt time.Time) *Transaction {
	return &Transaction{
		ID:        id,
		Status:    status,
		CreatedAt: createdAt,
	}
}

type Withdrawal struct {
	Transaction
	Amount float64 `json:"amount"`
}

func NewWithdrawal(transaction Transaction, amount float64) *Withdrawal {
	return &Withdrawal{
		Transaction: transaction,
		Amount:      amount,
	}
}

type Deposit struct {
	Transaction
	Amount float64 `json:"amount"`
}

func NewDeposit(transaction Transaction, amount float64) *Deposit {
	return &Deposit{
		Transaction: transaction,
		Amount:      amount,
	}
}

type CheckDeposit struct {
	Deposit
	CheckNumber string `json:"check_number"`
	BankCode    string `json:"bank_code"`
}

func NewCheckDeposit(checkNumber, bankCode string, deposit Deposit) *CheckDeposit {
	return &CheckDeposit{
		Deposit:     deposit,
		CheckNumber: checkNumber,
		BankCode:    bankCode,
	}
}

type CashDeposit struct {
	Deposit
	CashDepositLimit float64 `json:"cash_deposit_limit"`
}

func NewCashDeposit(cashDepositLimit float64, deposit Deposit) *CashDeposit {
	return &CashDeposit{
		Deposit:          deposit,
		CashDepositLimit: cashDepositLimit,
	}
}

type BalanceInquiry struct {
	Transaction
	AccountId string `json:"account_id"`
}

func NewBalanceInquiry(accountId string, transaction Transaction) *BalanceInquiry {
	return &BalanceInquiry{
		Transaction: transaction,
		AccountId:   accountId,
	}
}

type Transfer struct {
	Transaction
	FromAccountNumber string  `json:"from_account_number"`
	ToAccountNumber   string  `json:"to_account_number"`
	Amount            float64 `json:"amount"`
}

func NewTransfer(fromAccountNumber, toAccountNumber string, amount float64, transaction Transaction) *Transfer {
	return &Transfer{
		Transaction:       transaction,
		FromAccountNumber: fromAccountNumber,
		ToAccountNumber:   toAccountNumber,
		Amount:            amount,
	}
}
