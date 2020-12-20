package domain

import (
	"banking/errs"
	"banking/logger"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// CustomerRepositoryDb ..
type CustomerRepositoryDb struct {
	client *sqlx.DB
}

// FindAll ..
func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {

	var err error
	customers := make([]Customer, 0)

	if status == "" {
		findAllSQL := "SELECT * FROM customers"
		err = d.client.Select(&customers, findAllSQL)
	} else {
		findAllSQL := "SELECT * FROM customers WHERE status = ?"
		err = d.client.Select(&customers, findAllSQL, status)
	}

	if err != nil {
		logger.Error("Error while querying customers table: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected Error")
	}

	return customers, nil
}

// FindByID ..
func (d CustomerRepositoryDb) FindByID(id string) (*Customer, *errs.AppError) {

	customerSQL := "SELECT * FROM customers WHERE id = ?"
	var c Customer
	err := d.client.Get(&c, customerSQL, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer Not Found")
		}
		logger.Error("Error while querying customers table: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected Error")
	}
	return &c, nil
}

// NewCustomerRepositoryDb ..
func NewCustomerRepositoryDb() CustomerRepositoryDb {
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
	return CustomerRepositoryDb{client: client}
}
