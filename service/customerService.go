package service

import "banking/domain"

// CustomerService ..
type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
}

// DefaultCustomerService ..
type DefaultCustomerService struct {
	repo domain.CustomerRespository
}

// GetAllCustomer ..
func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

// NewCustomerService ..
func NewCustomerService(repository domain.CustomerRespository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
