package service

import (
	"abank/domain"
	"abank/dto"
	"abank/errs"
)

type CustomerService interface {
	GetAllCustomers(status string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustmerBy(id string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func NewDefaultCustomerService(repo domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repo}
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]dto.CustomerResponse, *errs.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}

	customers, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}

	dtos := make([]dto.CustomerResponse, 0)
	for _, cucustomer := range customers {
		dtos = append(dtos, cucustomer.ToDTO())
	}
	return dtos, nil
}

func (s DefaultCustomerService) GetCustmerBy(id string) (*dto.CustomerResponse, *errs.AppError) {
	customer, err := s.repo.FindBy(id)
	if err != nil {
		return nil, err
	}

	dto := customer.ToDTO()
	return &dto, nil
}
