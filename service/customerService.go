package service

import (
	"capi/domain"
	"capi/errs"
)

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, *errs.AppErr)
	GetCustomerByID(string) (*domain.Customer, *errs.AppErr)
}

type DefaultCustomerService struct {
	repository domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, *errs.AppErr) {
	// add process here
	return s.repository.FindAll()
}

func (s DefaultCustomerService) GetCustomerByID(customerID string) (*domain.Customer, *errs.AppErr) {
	return s.repository.FindByID(customerID)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository: repository}
}
