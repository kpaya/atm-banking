package entity

type ATM struct {
	ID string `json:"id"`
	// Address Address `json:"address"`
}

func NewATM(id string) *ATM {
	return &ATM{
		ID: id,
	}
}
