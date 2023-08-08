package value_object

type Address struct {
	ID      string `json:"id" gorm:"primaryKey,default:gen_random_uuid()"`
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	Zipcode string `json:"zipcode"`
	Country string `json:"country"`
}

func NewAddress(street, city, state, zipcode, country string) *Address {
	return &Address{
		Street:  street,
		City:    city,
		State:   state,
		Zipcode: zipcode,
		Country: country,
	}
}
