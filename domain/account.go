package domain

import (
	"banking/dto"
	"banking/errs"
)

// Account ..
type Account struct {
	AccountID   string `db:"id"`
	CustomerID  string `db:"customer_id"`
	OpeningDate string `db:"opening_date"`
	AccountType string `db:"account_type"`
	Amount      int64  `db:"amount"`
	Status      string `db:"status"`
}

// AccountRepository ..
type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
	SaveTransaction(transaction Transaction) (*Transaction, *errs.AppError)
	FindByID(accountID string) (*Account, *errs.AppError)
}

// ToNewAccountResponseDto ..
func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{AccountID: a.AccountID}
}

// CanWithdraw ..
func (a Account) CanWithdraw(amount int64) bool {
	if a.Amount >= amount {
		return true
	}
	return false
}
