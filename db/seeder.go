package db

import (
	"database/sql"
	"log"
)

var events = []struct{
	ID int
	Name, StartDate, StartTime string
}{
	{
		ID: 1,
		Name: "Killer whale entertainmant show",
		StartDate: "2021-11-19",
		StartTime: "18:45:00",
	},
	{
		ID: 2,
		Name: "Rock'n Roll",
		StartDate: "2021-11-19",
		StartTime: "18:45:00",
	},
	{
		ID: 3,
		Name: "Art ceremony",
		StartDate: "2021-09-19",
		StartTime: "18:45:00",
	},
	{
		ID: 4,
		Name: "October fest",
		StartDate: "2021-11-19",
		StartTime: "18:45:00",
	},
	{
		ID: 5,
		Name: "World Cup Final",
		StartDate: "2022-09-19",
		StartTime: "18:45:00",
	},

}


var ticketTypes = []struct{
	ID, SellingOption int
	Name, Currency string
	Price float64
}{
	{
		ID: 1,
		Name: "normal",
		Price: 19.99,
		SellingOption: 0,
		Currency: "EURO",
	},
	{
		ID: 2,
		Name: "golder",
		Price: 39.99,
		SellingOption: 1,
		Currency: "EURO",
	},
	{
		ID: 3,
		Name: "even",
		Price: 29.99,
		SellingOption: 2,
		Currency: "EURO",
	},
	{
		ID: 4,
		Name: "one",
		Price: 9.99,
		SellingOption: 3,
		Currency: "EURO",
	},
}

//Seed seed database inital data
func Seed(db *sql.DB)  {
	seedEvents(db)
	seedEventTickets(db)
	seedTicketInitialQuantities(db)
}

func seedEvents(db *sql.DB) {
	for _, v := range events {
		query := `INSERT IGNORE events(id, name, start_date, start_time) VALUES(?,?,?,?)`
		_, err := db.Exec(query, v.ID, v.Name, v.StartDate, v.StartTime)

		if err != nil {
			log.Println(err)
		}
	}
}

func seedTicketInitialQuantities(db *sql.DB) {
	eventIds := getEventIds(db)
	ticketTypeIds := getTicketTypeIds(db)

	for _, e := range eventIds {
		for i, t := range ticketTypeIds{
			quantity := i*10 + 10
			query := `INSERT IGNORE event_tickets(event_id, ticket_type_id, quantity) VALUES(?,?,?)`
			_, err := db.Exec(query, e, t, quantity)

			if err != nil {
				log.Println(err)
			}
		}
	}
}

func seedEventTickets(db *sql.DB) {
	for _, v := range ticketTypes {
		query := `INSERT IGNORE ticket_types(id, name, price, currency, selling_option) VALUES(?,?,?,?,?)`
		_, err := db.Exec(query, v.ID, v.Name, v.Price, v.Currency, v.SellingOption)

		if err != nil {
			log.Println(err)
		}
	}
}

func getEventIds(db *sql.DB) []int {
	var events []int
	query := `SELECT id FROM events`

	rows, err := db.Query(query)

	if err != nil {
		log.Println(err)
		return events
	}

	defer rows.Close()

	for rows.Next() {
		var id int

		if err := rows.Scan(&id); err != nil {
			log.Println(err)
			continue
		}

		events = append(events, id)
	}

	return events
}

func getTicketTypeIds(db *sql.DB) []int {
	var tickets []int
	query := `SELECT id FROM ticket_types`

	rows, err := db.Query(query)

	if err != nil {
		log.Println(err)
		return tickets
	}

	defer rows.Close()

	for rows.Next() {
		var id int

		if err := rows.Scan(&id); err != nil {
			log.Println(err)
			continue
		}

		tickets = append(tickets, id)
	}

	return tickets
}