package domain

import (
	"banking/dto"
	"banking/errs"
)

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

func (c Customer) asStatusText() string {
	status := "inactive"
	if c.Status == "1" {
		status = "active"
	}
	return status
}

// ToDto ..
func (c Customer) ToDto() dto.CustomerResponse {
	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateofBirth: c.DateofBirth,
		Status:      c.asStatusText(),
	}
}
