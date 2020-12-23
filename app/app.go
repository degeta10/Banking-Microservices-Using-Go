package app

import (
	"banking/domain"
	"banking/service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

// Start ..
func Start() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := mux.NewRouter()

	client := getDbClient()

	customerRepositoryDb := domain.NewCustomerRepositoryDb(client)
	accountRepositoryDb := domain.NewAccountRepositoryDb(client)

	// Handlers
	ch := CustomerHandlers{service: service.NewCustomerService(customerRepositoryDb)}
	ah := AccountHandlers{service: service.NewAccountService(accountRepositoryDb)}

	// Routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customer/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customer/{customer_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)

	// Start Server
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}

// createCustomer ..
func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post request")
}

// getCustomer ..
func getCustomer(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	fmt.Fprint(w, v["customer_id"])
}

// getDbClient ..
func getDbClient() *sqlx.DB {
	driver := os.Getenv("DATABASE_DRIVER")
	dbname := os.Getenv("DATABASE_NAME")
	dbhost := os.Getenv("DATABASE_HOST")
	dbport := os.Getenv("DATABASE_PORT")
	dbuser := os.Getenv("DATABASE_USER")
	dbpassword := os.Getenv("DATABASE_PASSWORD")
	client, err := sqlx.Open(fmt.Sprintf("%s", driver), fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbuser, dbpassword, dbhost, dbport, dbname))
	if err != nil {
		panic(err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}

// writeResponse ..
func writeResponse(r *http.Request, w http.ResponseWriter, code int, data interface{}) {
	if r.Header.Get("Content-Type") == "application/xml" {
		// xml response
		writeXMLResponse(w, code, data)
	} else {
		// json response
		writeJSONResponse(w, code, data)
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
