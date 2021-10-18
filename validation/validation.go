package validation

import "github.com/asaskevich/govalidator"

//CreateBookingRequest validate booking creation
type CreateBookingRequest struct {
	UserID int `json:"user_id"`
	EventId int `json:"event_id" valid:"required"`
	TicketTypeId int `json:"ticket_type_id" valid:"required"`
	Quantity int `json:"quantity" valid:"required"`
}

func (c *CreateBookingRequest) Validate() (bool, error) {
	return govalidator.ValidateStruct(c)
}

