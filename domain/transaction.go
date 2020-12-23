package domain

import (
	"banking/dto"
	"banking/errs"
)

// Transaction ..
type Transaction struct {
	TransactionID   string `db:"id"`
	AccountID       string `db:"account_id"`
	TransactionDate string `db:"date"`
	TransactionType string `db:"transaction_type"`
	Amount          int64  `db:"amount"`
}

// TransactionRepository ..
type TransactionRepository interface {
	Save(Transaction) (*Transaction, *errs.AppError)
}

// ToNewTransactionResponseDto ..
func (t Transaction) ToNewTransactionResponseDto() dto.NewTransactionResponse {
	return dto.NewTransactionResponse{AccountID: t.AccountID, TransactionID: t.TransactionID, Balance: t.Amount}
}

// isWithdrawal ..
func (t Transaction) isWithdrawal() bool {
	if t.TransactionType == "withdrawal" {
		return true
	}
	return false
}
