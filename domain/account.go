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
}

// ToNewAccountResponseDto ..
func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{AccountID: a.AccountID}
}
