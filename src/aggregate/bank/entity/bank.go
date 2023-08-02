package entity

type Bank struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

func NewBank(name, code string) *Bank {
	return &Bank{
		Name: name,
		Code: code,
	}
}
