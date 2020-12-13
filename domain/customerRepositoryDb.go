package domain

import (
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
func (d CustomerRepositoryDb) FindAll() ([]Customer, error) {

	findAllSql := "SELECT * FROM customers"
	rows, err := d.client.Query(findAllSql)
	if err != nil {
		log.Println("Error while querying customers table " + err.Error())
		return nil, err
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.DateofBirth, &c.Status)
		if err != nil {
			log.Println("Error while querying customers table " + err.Error())
			return nil, err
		}
		customers = append(customers, c)
	}

	return customers, nil
}

// NewCustomerRepositoryDb ..
func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sql.Open("mysql", "root:''@localhost/banking")
	if err != nil {
		panic(err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{client: client}
}
