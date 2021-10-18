package ticket

import (
	"database/sql"
)

type TicketRepository interface {
	GetById(id int) (*TicketType, error)
}

type ticketRepository struct {
	db *sql.DB
}

func NewTicketRepository(db *sql.DB) TicketRepository {
	return &ticketRepository{
		db: db,
	}
}

func (t *ticketRepository) GetById(id int) (*TicketType, error) {
	var tt TicketType

	query := `SELECT id, name, price, currency, selling_option FROM ticket_types WHERE id = ? LIMIT 1`

	stmt, err := t.db.Prepare(query)

	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(id)

	if err := row.Scan(&tt.ID, &tt.Name, &tt.Price, &tt.Currency, &tt.SellingOption); err != nil {
		return nil, err
	}

	return &tt, nil
}