package service

import (
	"banking/domain"
	"banking/errs"
)

// CustomerService ..
type CustomerService interface {
	GetAllCustomer(string) ([]domain.Customer, *errs.AppError)
	GetCustomer(string) (*domain.Customer, *errs.AppError)
}

// DefaultCustomerService ..
type DefaultCustomerService struct {
	repo domain.CustomerRespository
}

// GetAllCustomer ..
func (s DefaultCustomerService) GetAllCustomer(status string) ([]domain.Customer, *errs.AppError) {
	if status == "active" {
		return s.repo.FindAll("1")
	} else if status == "inactive" {
		return s.repo.FindAll("0")
	} else {
		return s.repo.FindAll("")
	}
}

// GetCustomer ..
func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppError) {
	return s.repo.FindByID(id)
}

// NewCustomerService ..
func NewCustomerService(repository domain.CustomerRespository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
