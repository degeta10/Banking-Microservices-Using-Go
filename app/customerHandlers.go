package app

import (
	"banking/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// CustomerHandlers ..
type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {

	c, err := ch.service.GetAllCustomer()

	if err != nil {
		if r.Header.Get("Content-Type") == "application/xml" {
			// xml response
			writeXMLResponse(w, err.Code, err.AsMessage())
		} else {
			// json response
			writeJSONResponse(w, err.Code, err.AsMessage())
		}
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		// xml response
		writeXMLResponse(w, http.StatusOK, c)
	} else {
		// json response
		writeJSONResponse(w, http.StatusOK, c)
	}
}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["customer_id"]
	c, err := ch.service.GetCustomer(id)

	if err != nil {
		if r.Header.Get("Content-Type") == "application/xml" {
			// xml response
			writeXMLResponse(w, err.Code, err.AsMessage())
		} else {
			// json response
			writeJSONResponse(w, err.Code, err.AsMessage())
		}
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		// xml response
		writeXMLResponse(w, http.StatusOK, c)
	} else {
		// json response
		writeJSONResponse(w, http.StatusOK, c)
	}
}

// writeXMLResponse ..
func writeXMLResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/xml")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}

// writeJSONResponse ..
func writeJSONResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}
