package customer

type CustomerStatus string

const (
	ACTIVE      CustomerStatus = "Active"
	BLOCKED     CustomerStatus = "Blocked"
	BANNED      CustomerStatus = "Banned"
	COMPROMISED CustomerStatus = "Compromised"
	ARCHIVED    CustomerStatus = "Archived"
	CLOSED      CustomerStatus = "Closed"
	UNKNOWN     CustomerStatus = "Unknown"
)

type Customer struct {
	Name   string         `json:"name"`
	Email  string         `json:"email"`
	Phone  string         `json:"phone"`
	Status CustomerStatus `json:"status"`
	// Card Card
	// Account BankingAccount
	// Address Address
}

func NewCustomer(name, email, phone string) *Customer {
	return &Customer{
		Name:   name,
		Email:  email,
		Phone:  phone,
		Status: ACTIVE,
	}
}
