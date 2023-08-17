package dto

type InputAccessAccountDTO struct {
	AccountNumber string `json:"account_number"`
	Pin           string `json:"pin"`
}

type OutputAccessAccountDTO struct {
	Key string `json:"key"`
}
