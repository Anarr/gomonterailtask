package handler

import (
	"github.com/Anarr/gomonterailtask/usecase"
	"github.com/Anarr/gomonterailtask/utils"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

//BookingHandler handle create booking
func BookingHandler(service usecase.BookingUseCase) httprouter.Handle {
	return func(w http.ResponseWriter, r * http.Request, _ httprouter.Params) {
		b, err := service.Book(r)

		if err != nil {
			utils.Error(w, 422,err.Error())
			return
		}

		utils.Success(w, b)
		return
	}
}

//BookingConfirmationHandler handle booking confirmation
func BookingConfirmationHandler(service usecase.BookingUseCase) httprouter.Handle {
	return func(w http.ResponseWriter, r * http.Request, params httprouter.Params) {

		id, _ := strconv.Atoi(params.ByName("id"))
		b, err := service.Confirm(r, id)

		if err != nil {
			utils.Error(w, 422, err.Error())
			return
		}

		utils.Success(w, b)
		return
	}
}

