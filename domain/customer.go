package domain

import "banking/errs"

// Customer ..
type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateofBirth string
	Status      string
}

// CustomerRespository ..
type CustomerRespository interface {
	FindAll(status string) ([]Customer, *errs.AppError)
	FindByID(string) (*Customer, *errs.AppError)
}
