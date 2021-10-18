package usecase

import (
	"encoding/json"
	"github.com/Anarr/gomonterailtask/repository/booking"
	"github.com/Anarr/gomonterailtask/repository/event"
	"github.com/Anarr/gomonterailtask/repository/ticket"
	"github.com/Anarr/gomonterailtask/validation"
	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"net/http"
)

type BookingUseCase interface {
	Book(r *http.Request) (*booking.Booking, error)
	Confirm(r *http.Request, id int) (*booking.Booking, error)
}

type bookingService struct {
	bookRepository booking.BookingRepository
	ticketRepository ticket.TicketRepository
	eventRepository event.EventRepository
}

//Tickets retrieve all tickets
func (t *bookingService) Book(r *http.Request)  (*booking.Booking, error) {

	var createBookingRequest validation.CreateBookingRequest
	//pass userid manuall
	authId := r.Header.Get("user_id")

	userID,err := govalidator.ToInt(authId)

	if err != nil {
		return nil, err
	}

	if err := json.NewDecoder(r.Body).Decode(&createBookingRequest); err != nil {
		return nil, err
	}

	if _, err := createBookingRequest.Validate(); err != nil {
		return nil, err
	}

	if err := t.eventRepository.IsExpired(createBookingRequest.EventId); err != nil {
		return nil, err
	}

	createBookingRequest.UserID = int(userID)

	transactionId, _ := uuid.NewUUID()

	data := booking.Booking{
		TransactionID: transactionId.String(),
		EventID:       createBookingRequest.EventId,
		TicketTypeID:  createBookingRequest.TicketTypeId,
		Quantity:      createBookingRequest.Quantity,
		UserID: createBookingRequest.UserID,
	}

	ticketType, err := t.ticketRepository.GetById(data.TicketTypeID)

	if err != nil {
		return nil, err
	}

	eventTicket, err := t.eventRepository.Ticket(data.EventID, data.TicketTypeID)

	if err != nil {
		return nil, err
	}

	if err := t.bookRepository.IsValidQuantity(ticketType, &data, eventTicket); err != nil {
		return nil, err
	}

	return t.bookRepository.Book(data)
}

//Tickets retrieve all tickets
func (t *bookingService) Confirm(r *http.Request, id int) (*booking.Booking, error)  {
	authId := r.Header.Get("user_id")
	userID,err := govalidator.ToInt(authId)

	if err != nil {
		return nil, err
	}

	b := booking.Booking{
		UserID: int(userID),
		ID: id,
	}

	return t.bookRepository.Confirm(b)
}

//NewBookingUseCase create new BookingUseCase instance
func NewBookingUseCase(repository booking.BookingRepository, ticketRepository ticket.TicketRepository, eventRepository event.EventRepository) BookingUseCase {
	return &bookingService{
		bookRepository: repository,
		ticketRepository: ticketRepository,
		eventRepository: eventRepository,
	}
}
