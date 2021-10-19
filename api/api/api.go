package api

import (
	"github.com/Anarr/gomonterailtask/api/handler"
	"github.com/Anarr/gomonterailtask/db"
	"github.com/Anarr/gomonterailtask/repository/booking"
	"github.com/Anarr/gomonterailtask/repository/event"
	"github.com/Anarr/gomonterailtask/repository/ticket"
	"github.com/Anarr/gomonterailtask/usecase"
	"github.com/julienschmidt/httprouter"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

//Run run application
func Run(config *viper.Viper) {
	//initalize db connectionp
	dbConf := &db.MysqlDb{
		Host: config.GetString("database.host"),
		Username: config.GetString("database.user"),
		Pass: config.GetString("database.password"),
		Name: config.GetString("database.name"),
		Port: config.GetInt("database.port"),
	}

	database, err := db.New(dbConf)

	if err != nil {
		log.Fatal(err)
	}

	db.Seed(database)
	//regiser router
	router := httprouter.New()
	//events
	eventRepository := event.NewEventRepository(database)
	bookingRepository := booking.NewBookingRepository(database)
	ticketRepository := ticket.NewTicketRepository(database)
	eventUseCase := usecase.NewEventUseCase(eventRepository)
	router.GET("/api/events", handler.EventHandler(eventUseCase))
	router.GET("/api/events/:id/tickets", handler.EventAvailableTicketHandler(eventUseCase))

	//bookings
	bookingUseCase := usecase.NewBookingUseCase(bookingRepository, ticketRepository, eventRepository)
	router.POST("/api/bookings", handler.BookingHandler(bookingUseCase))
	router.PATCH("/api/bookings/:id", handler.BookingConfirmationHandler(bookingUseCase))


	err = http.ListenAndServe(":8001", router); if err != nil {
		log.Fatal(err)
	}
}
