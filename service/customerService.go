package service

import (
	"banking/domain"
	"banking/errs"
)

// CustomerService ..
type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, *errs.AppError)
	GetCustomer(string) (*domain.Customer, *errs.AppError)
}

// DefaultCustomerService ..
type DefaultCustomerService struct {
	repo domain.CustomerRespository
}

// GetAllCustomer ..
func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, *errs.AppError) {
	return s.repo.FindAll()
}

// GetCustomer ..
func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppError) {
	return s.repo.FindByID(id)
}

// NewCustomerService ..
func NewCustomerService(repository domain.CustomerRespository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
