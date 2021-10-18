package event

//Event store event model
type Event struct {
	ID int `json:"id"`
	Name string `json:"name"`
	StartDate string `json:"start_date"`
	StartTime string `json:"start_time"`
	CreatedAt string `json:"-"`
}

//EventTicket store event tickets model
type EventTicket struct {
	ID int `json:"id"`
	EventId int `json:"event_id"`
	TicketTypeId int `json:"ticket_type_id"`
	Quantity int `json:"quantity"`
	ConfirmedBookingCount int `json:"confirmed_booking_count"`
}

//GetAvailableQuantity calculate available quantity count consider exists confirmed bookings
func (et *EventTicket) GetAvailableQuantity() int {
	return et.Quantity - et.ConfirmedBookingCount
}
