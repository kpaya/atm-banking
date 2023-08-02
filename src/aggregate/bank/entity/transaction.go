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

func (w *Withdrawal) GetAmount() float64 {
	return w.Amount
}

type Deposit struct {
	Transaction
	Amount float64 `json:"amount"`
}

func (d *Deposit) GetAmount() float64 {
	return d.Amount
}
