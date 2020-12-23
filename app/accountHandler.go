package app

import (
	"banking/dto"
	"banking/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// AccountHandlers ..
type AccountHandlers struct {
	service service.AccountService
}

// NewAccount ..
func (h *AccountHandlers) NewAccount(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	cid := vars["customer_id"]

	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		writeResponse(r, w, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerID = cid
		account, appError := h.service.NewAccount(request)
		if appError != nil {
			writeResponse(r, w, appError.Code, appError.Message)
		}
		writeResponse(r, w, http.StatusCreated, account)
	}
}

// NewTransaction ..
func (h *AccountHandlers) NewTransaction(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	cid := vars["customer_id"]
	aid := vars["account_id"]

	var request dto.NewTransactionRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		writeResponse(r, w, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerID = cid
		request.AccountID = aid
		account, appError := h.service.NewTransaction(request)
		if appError != nil {
			writeResponse(r, w, appError.Code, appError.Message)
		}
		writeResponse(r, w, http.StatusCreated, account)
	}
}
