package customer

type account struct {
	AccountNumber    string  `json:"account_number"`
	TotalBalance     float64 `json:"total_balance"`
	AvaliableBalance float64 `json:"avaliable_balance"`
}

type SavingsAccount struct {
	account
	WithdrawLimit float64 `json:"withdraw_limit"`
}

func NewSavingAccount(accountNumber string, totalBalance, avaliableBalance, withdrawLimit float64) *SavingsAccount {
	return &SavingsAccount{
		account: account{
			AccountNumber:    accountNumber,
			TotalBalance:     totalBalance,
			AvaliableBalance: avaliableBalance,
		},
		WithdrawLimit: withdrawLimit,
	}
}

type CheckingAccount struct {
	account
	DebitCardNumber string `json:"debit_card_number"`
}

func NewCheckingAccount(accountNumber string, totalBalance, avaliableBalance float64, debitCardNumber string) *CheckingAccount {
	return &CheckingAccount{
		account: account{
			AccountNumber:    accountNumber,
			TotalBalance:     totalBalance,
			AvaliableBalance: avaliableBalance,
		},
		DebitCardNumber: debitCardNumber,
	}
}
