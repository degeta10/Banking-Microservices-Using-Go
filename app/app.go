package app

import (
	"banking/domain"
	"banking/service"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Start ..
func Start() {

	router := mux.NewRouter()

	// Handlers
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	// Routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customer/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	// Start Server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post request")
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	fmt.Fprint(w, v["customer_id"])
}
