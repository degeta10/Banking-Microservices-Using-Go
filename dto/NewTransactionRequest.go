package dto

import (
	"banking/errs"
	"strings"
)

// NewTransactionRequest ..
type NewTransactionRequest struct {
	CustomerID      string `json:"customer_id"`
	AccountID       string `json:"account_id"`
	TransactionType string `json:"transaction_type"`
	Amount          int64  `json:"amount"`
}

// Validate ..
func (r NewTransactionRequest) Validate() *errs.AppError {
	minAmount := 500
	if int(r.Amount) < 0 {
		return errs.NewValidationError("Amount cannot be less than zero")
	}
	if int(r.Amount) < minAmount {
		return errs.NewValidationError("Mininmum amount is 500")
	}
	if strings.ToLower(r.TransactionType) != "withdrawal" && strings.ToLower(r.TransactionType) != "deposit" {
		return errs.NewValidationError("Transaction type must be withdrawal or deposit")
	}
	return nil
}
