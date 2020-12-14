package domain

import "banking/errs"

// Customer ..
type Customer struct {
	Id          string `db:"id"`
	Name        string `db:"name"`
	City        string `db:"city"`
	Zipcode     string `db:"zipcode"`
	DateofBirth string `db:"date_of_birth"`
	Status      string `db:"status"`
}

// CustomerRespository ..
type CustomerRespository interface {
	FindAll(status string) ([]Customer, *errs.AppError)
	FindByID(string) (*Customer, *errs.AppError)
}
