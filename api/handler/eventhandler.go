package handler

import (
	"github.com/Anarr/gomonterailtask/usecase"
	"github.com/Anarr/gomonterailtask/utils"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"strconv"
)

//EventHandler retrive tickets list
func EventHandler(service usecase.EventUseCase) httprouter.Handle {
	return func(w http.ResponseWriter, r * http.Request, _ httprouter.Params) {
		tickets, err := service.All()

		if err != nil {
			utils.Error(w, 422, "Can not fetch events")
			return
		}

		utils.Success(w, tickets)
		return
	}
}

//EventAvailableTicketHandler retrive tickets list
func EventAvailableTicketHandler(service usecase.EventUseCase) httprouter.Handle {
	return func(w http.ResponseWriter, r * http.Request, ps httprouter.Params) {
		param := ps.ByName("id")
		id, err := strconv.Atoi(param)
		if err != nil {
			utils.Error(w, 422, "Can not fetch events")
			return
		}

		tickets, err := service.AvailableTickets(id)

		if err != nil {
			utils.Error(w, 422, "Can not fetch events")
			return
		}

		utils.Success(w, tickets)
		return
	}
}

