package ticket

const NONE = 0
const Together = 1
const Even = 2
const One = 3

//TicketType
type TicketType struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
	Currency string `json:"currency"`
	SellingOption int `json:"selling_option"`
}

