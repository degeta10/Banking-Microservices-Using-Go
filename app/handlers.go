package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

// Customer ...
type Customer struct {
	Name    string `json:"name" xml:"name" `
	City    string `json:"city"  xml:"city"`
	Zipcode string `json:"zipcode"  xml:"zipcode"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world!")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	c := []Customer{
		{Name: "A", City: "A", Zipcode: "A"},
		{Name: "B", City: "B", Zipcode: "B"},
		{Name: "C", City: "C", Zipcode: "C"},
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
