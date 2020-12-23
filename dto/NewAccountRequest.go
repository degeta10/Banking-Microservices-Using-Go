package dto

import (
	"banking/errs"
	"strings"
)

// NewAccountRequest ..
type NewAccountRequest struct {
	CustomerID  string `json:"customer_id"`
	AccountType string `json:"account_type"`
	Amount      int64  `json:"amount"`
}

// Validate ..
func (r NewAccountRequest) Validate() *errs.AppError {
	minAmount := 1000
	if int(r.Amount) < minAmount {
		return errs.NewValidationError("Mininmum amount is 1000")
	}
	if strings.ToLower(r.AccountType) != "saving" && strings.ToLower(r.AccountType) != "checking" {
		return errs.NewValidationError("Account type must be saving or checking")
	}
	return nil
}
