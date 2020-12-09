package app

import (
	"banking/service"
	"encoding/json"
	"encoding/xml"
	"net/http"
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
