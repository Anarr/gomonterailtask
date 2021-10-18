package event

import (
	"database/sql"
	"errors"
	"log"
)

type EventRepository interface {
	All() ([]Event, error)
	Tickets(id int) ([]EventTicket, error)
	Ticket(id, ticketTypeId int) (*EventTicket, error)
	IsExpired(id int) error
}

type eventRepository struct {
	db *sql.DB
}

func NewEventRepository(db *sql.DB) EventRepository {
	return &eventRepository{
		db: db,
	}
}

//All retrieve all tickets
func (t *eventRepository) All() ([]Event, error) {

	result := []Event{}
	query := `SELECT id, name, start_date, start_time FROM events`

	stmt, err := t.db.Prepare(query)

	if err != nil {
		log.Println(err)
		return result, err
	}

	rows, err := stmt.Query()

	defer rows.Close()

	if err != nil {
		log.Println(err)
		return result, err
	}

	for rows.Next() {
		var e Event

		if err = rows.Scan(&e.ID, &e.Name, &e.StartDate, &e.StartTime); err != nil {
			continue
		}

		result = append(result, e)
	}

	return result, nil
}

//Tickets retrieve event all available tickets
func (t *eventRepository) Tickets(id int) ([]EventTicket, error) {
	result := []EventTicket{}

	query := `SELECT et.id, et.event_id, et.ticket_type_id, 
       			et.quantity,
                IF(sum(b.quantity)>0, sum(b.quantity), 0) as selled_count
				FROM event_tickets et
                left join bookings b ON b.event_id = et.event_id and b.ticket_type_id = et.ticket_type_id and b.status = 1
				WHERE et.event_id = ?
                group by et.id`

	stmt, err := t.db.Prepare(query)

	if err != nil {
		log.Println(err)
		return result, err
	}

	rows, err := stmt.Query(id)

	defer rows.Close()

	if err != nil {
		log.Println(err)
		return result, err
	}

	for rows.Next() {
		var e EventTicket

		if err = rows.Scan(&e.ID, &e.EventId, &e.TicketTypeId, &e.Quantity, &e.ConfirmedBookingCount); err != nil {
			continue
		}

		result = append(result, e)
	}

	return result, nil
}

//Ticket get single ticket by event_id and ticket_type_id
func (t *eventRepository) Ticket(id, ticketTypeId int) (*EventTicket, error) {
	var et EventTicket

	query := `SELECT et.id, et.event_id, et.ticket_type_id, et.quantity,
       	IF(sum(b.quantity)>0, sum(b.quantity), 0) as selled_count
		FROM event_tickets et
		left join bookings b ON b.event_id = et.event_id and b.ticket_type_id = et.ticket_type_id and b.status = 1
		WHERE et.event_id = ?  and et.ticket_type_id = ? group by et.id LIMIT 1`

	stmt, err := t.db.Prepare(query)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	row := stmt.QueryRow(id, ticketTypeId)

	if err = row.Scan(&et.ID, &et.EventId, &et.TicketTypeId, &et.Quantity, &et.ConfirmedBookingCount); err != nil {
		log.Println(err)
		return nil, err
	}


	return &et, nil
}

func (t *eventRepository) IsExpired(id int) error  {
	var e Event

	query := `SELECT id  
				FROM events 
				WHERE id = ? AND DATE(NOW()) < start_date OR (DATE(NOW()) = start_date AND date_add(CURRENT_TIME, INTERVAL 15 MINUTE) < start_time)
				LIMIT 1`

	stmt, err := t.db.Prepare(query)

	if err != nil {
		log.Println(err)
		return errors.New("event booking date expired")
	}

	row := stmt.QueryRow(id)

	if err = row.Scan(&e.ID); err != nil {
		log.Println(err)
		return errors.New("event booking date expired")
	}

	if e.ID == 0 {
		return errors.New("event booking date expired")
	}

	return nil
}