package domain

import (
	"banking/errs"
	"banking/logger"
	"database/sql"
	"strconv"

	"github.com/jmoiron/sqlx"
)

// AccountRepositoryDb ..
type AccountRepositoryDb struct {
	client *sqlx.DB
}

// FindByID ..
func (d AccountRepositoryDb) FindByID(id string) (*Account, *errs.AppError) {

	accountSQL := "SELECT * FROM accounts WHERE id = ?"
	var a Account
	err := d.client.Get(&a, accountSQL, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Account Not Found")
		}
		logger.Error("Error while querying accounts table: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected Error")
	}
	return &a, nil
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

// SaveTransaction ..
func (d AccountRepositoryDb) SaveTransaction(t Transaction) (*Transaction, *errs.AppError) {

	tx, err := d.client.Begin()
	if err != nil {
		logger.Error("Error while starting transaction block: " + err.Error())
		return nil, errs.NewUnexpectedError("Transaction block failed")
	}

	sqlInsert := "INSERT INTO transactions (account_id,date,amount,transaction_type) values (?,?,?,?)"
	result, err := tx.Exec(sqlInsert, t.AccountID, t.TransactionDate, t.Amount, t.TransactionType)
	if err != nil {
		tx.Rollback()
		logger.Error("Error while creation transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Transaction Creation Error")
	}

	if t.isWithdrawal() {
		_, err = tx.Exec("UPDATE accounts SET amount = amount - ? WHERE id = ?", t.Amount, t.AccountID)
	} else {
		_, err = tx.Exec("UPDATE accounts SET amount = amount + ? WHERE id = ?", t.Amount, t.AccountID)
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Error while updating account: " + err.Error())
		return nil, errs.NewUnexpectedError("Account Updating Error")
	}

	id, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		logger.Error("Error while getting account: " + err.Error())
		return nil, errs.NewUnexpectedError("Account Fetching Error")
	}

	account, appErr := d.FindByID(t.AccountID)
	if err != nil {
		return nil, appErr
	}

	t.TransactionID = strconv.FormatInt(id, 10)
	t.Amount = account.Amount
	return &t, nil
}

// NewAccountRepositoryDb ..
func NewAccountRepositoryDb(client *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{client: client}
}
