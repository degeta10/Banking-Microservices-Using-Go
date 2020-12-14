package domain

import (
	"banking/errs"
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// CustomerRepositoryDb ..
type CustomerRepositoryDb struct {
	client *sql.DB
}

// FindAll ..
func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {

	var rows *sql.Rows
	var err error

	if status == "" {
		findAllSQL := "SELECT * FROM customers"
		rows, err = d.client.Query(findAllSQL)
	} else {
		findAllSQL := "SELECT * FROM customers WHERE status = ?"
		rows, err = d.client.Query(findAllSQL, status)
	}

	if err != nil {
		log.Println("Error while querying customers table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected Error")
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.DateofBirth, &c.City, &c.Zipcode, &c.Status)
		if err != nil {
			log.Println("Error while querying customers table " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected Error")
		}
		customers = append(customers, c)
	}

	return customers, nil
}

// FindByID ..
func (d CustomerRepositoryDb) FindByID(id string) (*Customer, *errs.AppError) {

	customerSQL := "SELECT * FROM customers WHERE id = ?"
	row := d.client.QueryRow(customerSQL, id)

	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.DateofBirth, &c.City, &c.Zipcode, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer Not Found")
		}
		log.Println("Error while querying customers table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected Error")
	}
	return &c, nil
}

// NewCustomerRepositoryDb ..
func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sql.Open("mysql", "global:qwe123@/banking")
	if err != nil {
		panic(err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{client: client}
}
