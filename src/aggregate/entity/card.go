package customer

type Card struct {
	Number          string `json:"number"`
	CustomerName    string `json:"customer_name"`
	ExpirationMonth int    `json:"expiration_month"`
	ExpirationYear  int    `json:"expiration_year"`
	CVV             int    `json:"cvv"`
}

func NewCard(number, customerName string, expirationMonth, expirationYear, cvv int) *Card {
	return &Card{
		Number:          number,
		CustomerName:    customerName,
		ExpirationMonth: expirationMonth,
		ExpirationYear:  expirationYear,
		CVV:             cvv,
	}
}
