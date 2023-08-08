package customer

import "github.com/go-playground/validator/v10"

type Card struct {
	ID              string `gorm:"primaryKey,default:gen_random_uuid()" json:"id"`
	Number          string `json:"number"`
	CustomerName    string `json:"customer_name"`
	ExpirationMonth int    `json:"expiration_month"`
	ExpirationYear  int    `json:"expiration_year"`
	CVV             int    `json:"cvv"`
}

func NewCard(number, customerName string, expirationMonth, expirationYear, cvv int) *Card {
	card := new(Card)
	card.Number = number
	card.CustomerName = customerName
	card.ExpirationMonth = expirationMonth
	card.ExpirationYear = expirationYear
	card.CVV = cvv
	if err := card.Validate(); err != nil {
		return nil
	}
	return card
}

func (c *Card) Validate() error {
	return validator.New().Struct(c)
}
