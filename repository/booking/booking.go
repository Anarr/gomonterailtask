package booking

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Anarr/gomonterailtask/repository/event"
	"github.com/Anarr/gomonterailtask/repository/ticket"
	"log"
)

const (
	ErrTicketNotAvailable = "there is no enough tickets in store"
	ErrTogetherTicket = "we can only buy all the tickets of this type at once"
	ErrEvenTicket = "we sell tickets in pairs (eg. 2, 4, 6 ...)"
	ErrOneTicket = "we can only buy tickets in a quantity that will not leave only 1 ticket in the system after the transaction"
	ErrBookingNotUpdated = "can not update booking"
)

type BookingRepository interface {
	Book(b Booking) (*Booking, error)
	Confirm(b Booking) (*Booking, error)
	GetById(id int) (*Booking, error)
	IsValidQuantity(tt *ticket.TicketType, booking *Booking, et *event.EventTicket) error
}

type bookingRepository struct {
	db *sql.DB
}

func NewBookingRepository(db *sql.DB) BookingRepository {
	return &bookingRepository{
		db: db,
	}
}

//Book create new booking
func (t *bookingRepository) Book(b Booking) (*Booking, error) {

	query := `INSERT INTO bookings(transaction_id, user_id, event_id, ticket_type_id, quantity) 
			VALUES(?,?,?,?,?)`

	stmt, err := t.db.Prepare(query)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	res, err := stmt.Exec(b.TransactionID, b.UserID, b.EventID, b.TicketTypeID, b.Quantity)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	lastId, err := res.LastInsertId()

	if err != nil {
		log.Println(err)
	}

	b.ID = int(lastId)


	return &b, nil
}

//Confirm confirm booking
func (t *bookingRepository) Confirm(b Booking) (*Booking, error) {
	query := `UPDATE bookings SET status = 1 WHERE id = ? AND user_id = ? and status = 0`
	stmt, err := t.db.Prepare(query)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	res, err := stmt.Exec(b.ID, b.UserID)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	if affectedRows, _ := res.RowsAffected(); affectedRows == 0 {
		return nil, errors.New(ErrBookingNotUpdated)
	}

	return t.GetById(b.ID)
}

//GetById get single booking
func (t *bookingRepository) GetById(id int) (*Booking, error) {
	var b Booking

	query := `SELECT id, transaction_id, user_id, event_id, ticket_type_id, quantity, status FROM bookings WHERE id = ? LIMIT 1`

	stmt, err := t.db.Prepare(query)

	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(id)

	if err := row.Scan(&b.ID, &b.TransactionID, &b.UserID, &b.EventID, &b.TicketTypeID, &b.Quantity, &b.Status); err != nil {
		return nil, err
	}

	return &b, nil
}

//IsValidQuantity check given quantity is valid for given ticket type and given event
func (t *bookingRepository) IsValidQuantity(tt *ticket.TicketType, b *Booking, et *event.EventTicket) error {

	if et.GetAvailableQuantity() < b.Quantity {
		fmt.Println(et.GetAvailableQuantity(), b.Quantity)
		return errors.New(ErrTicketNotAvailable)
	}

	switch tt.SellingOption {
	case ticket.NONE:
		return nil
	case ticket.Together:
		if b.Quantity != et.GetAvailableQuantity() {
			return errors.New(ErrTogetherTicket)
		}
	case ticket.Even:
		if b.Quantity%2 != 0 {
			return errors.New(ErrEvenTicket)
		}
	case ticket.One:
		if et.GetAvailableQuantity() - b.Quantity <= 1 {
			return errors.New(ErrOneTicket)
		}
	}

	return nil
}
