package usecase

import (
	"github.com/Anarr/gomonterailtask/repository/event"
)

type EventUseCase interface {
	All() ([]event.Event, error)
	AvailableTickets(id int) ([]event.EventTicket, error)
}

type eventService struct {
	repository event.EventRepository
}

//Tickets retrieve all tickets
func (t *eventService) All() ([]event.Event, error) {
	return t.repository.All()
}

func (t *eventService) AvailableTickets(id int) ([]event.EventTicket, error)  {
	return t.repository.Tickets(id)
}

//NewEventUseCase create new EventUseCase instance
func NewEventUseCase(repository event.EventRepository) EventUseCase {
	return &eventService{
		repository: repository,
	}
}
