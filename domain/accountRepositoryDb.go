package domain

import (
	"banking/errs"
	"banking/logger"
	"strconv"

	"github.com/jmoiron/sqlx"
)

// AccountRepositoryDb ..
type AccountRepositoryDb struct {
	client *sqlx.DB
}

// Save ..
func (d AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {
	sqlInsert := "INSERT INTO accounts (customer_id,opening_date,account_type,amount,status) values (?,?,?,?,?)"
	result, err := d.client.Exec(sqlInsert, a.CustomerID, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while creation account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected Error")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting account: " + err.Error())
	}

	a.AccountID = strconv.FormatInt(id, 10)
	return &a, nil
}

// NewAccountRepositoryDb ..
func NewAccountRepositoryDb(client *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{client: client}
}
