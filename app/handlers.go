package app

import (
	"banking/service"
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/gorilla/mux"
)

// CustomerHandlers ..
type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {

	c, _ := ch.service.GetAllCustomer()

	// xml response
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(c)
	} else {
		// json response
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(c)
	}
}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["customer_id"]
	c, err := ch.service.GetCustomer(id)

	if err != nil {
		// xml response
		if r.Header.Get("Content-Type") == "application/xml" {
			w.Header().Add("Content-Type", "application/xml")
			w.WriteHeader(http.StatusNotFound)
			xml.NewEncoder(w).Encode(err)
		} else {
			// json response
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(err)
		}
	}

	// xml response
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(c)
	} else {
		// json response
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(c)
	}
}
