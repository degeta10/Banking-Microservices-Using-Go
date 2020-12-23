package app

import (
	"banking/service"
	"net/http"

	"github.com/gorilla/mux"
)

// CustomerHandlers ..
type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {

	status := r.URL.Query().Get("status")

	c, err := ch.service.GetAllCustomer(status)

	if err != nil {
		writeResponse(r, w, err.Code, err.AsMessage())
	}

	writeResponse(r, w, http.StatusOK, c)
}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["customer_id"]
	c, err := ch.service.GetCustomer(id)

	if err != nil {
		writeResponse(r, w, err.Code, err.AsMessage())
	}

	writeResponse(r, w, http.StatusOK, c)
}
