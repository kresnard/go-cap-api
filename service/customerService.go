package service

import (
	"capi/domain"
	"capi/dto"
	"capi/errs"
)

type CustomerService interface {
	GetAllCustomer(string) (*[]dto.CustomerResponse, *errs.AppErr)
	GetCustomerByID(string) (*dto.CustomerResponse, *errs.AppErr)
}

type DefaultCustomerService struct {
	repository domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer(status string) (*[]dto.CustomerResponse, *errs.AppErr) {
	// add process here
	customers, err := s.repository.FindAll(status)
	if err != nil {
		return nil, errs.NewUnexpectedError("error query data to customer table")
	}
	var customerResponse []dto.CustomerResponse
	for _, data := range customers {
		customerResponse = append(customerResponse, data.ToDTO())
	}
	return &customerResponse, nil

}

func (s DefaultCustomerService) GetCustomerByID(customerID string) (*dto.CustomerResponse, *errs.AppErr) {
	cust, err := s.repository.FindByID(customerID)
	if err != nil {
		return nil, err
	}

	response := cust.ToDTO()

	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository: repository}
}
