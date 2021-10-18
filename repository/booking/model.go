package booking

type Booking struct {
	ID int `json:"id"`
	TransactionID string `json:"transaction_id"`
	UserID int `json:"user_id"`
	EventID int `json:"event_id"`
	TicketTypeID int `json:"ticket_type_id"`
	Quantity int `json:"quantity"`
	Status int `json:"status"`
}
