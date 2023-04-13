package domain

import "abank/errs"

type CustomerRepositoryStub struct {
	customers []Customer
}

func NewCustomerRepositoryStub(status string) CustomerRepositoryStub {
	customers := []Customer{
		{
			Id:        "1001",
			FirstName: "Nazmul",
			LastName:  "Islam",
			City:      "Bankgok",
			Zipcode:   "10400",
			Status:    "Active",
		},
		{
			Id:        "1002",
			FirstName: "Sumi",
			LastName:  "Khatun",
			City:      "Bankgok",
			Zipcode:   "10400",
			Status:    "Active",
		},
		{
			Id:        "1003",
			FirstName: "Aminul",
			LastName:  "Islam",
			City:      "Dhaka",
			Zipcode:   "1216",
			Status:    "InActive",
		},
		{
			Id:        "1004",
			FirstName: "Raive",
			LastName:  "Khan",
			City:      "Dhaka",
			Zipcode:   "1212",
			Status:    "Active",
		},
	}
	return CustomerRepositoryStub{customers: customers}
}

func (r CustomerRepositoryStub) FindAll() ([]Customer, *errs.AppError) {
	return r.customers, nil
}
