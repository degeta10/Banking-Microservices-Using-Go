package service

import (
	"banking/domain"
	"banking/dto"
	"banking/errs"
)

// CustomerService ..
type CustomerService interface {
	GetAllCustomer(string) ([]domain.Customer, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
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
func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	resp := c.ToDto()
	return &resp, nil
}

// NewCustomerService ..
func NewCustomerService(repository domain.CustomerRespository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
