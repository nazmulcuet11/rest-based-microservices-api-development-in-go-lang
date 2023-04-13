package domain

import (
	"abank/dto"
	"abank/errs"
)

type Customer struct {
	Id          string `db:"customer_id"`
	FirstName   string `db:"first_name"`
	LastName    string `db:"last_name"`
	City        string `db:"city"`
	Zipcode     string `db:"zipcode"`
	DateOfBirth string `db:"date_of_birth"`
	Status      string `db:"status"`
}

func (c Customer) statusAsText() string {
	if c.Status == "0" {
		return "inactive"
	} else {
		return "active"
	}
}

func (c Customer) ToDTO() dto.CustomerResponse {
	return dto.CustomerResponse{
		Id:          c.Id,
		FirstName:   c.FirstName,
		LastName:    c.LastName,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateOfBirth: c.DateOfBirth,
		Status:      c.statusAsText(),
	}
}

type CustomerRepository interface {
	FindAll(status string) ([]Customer, *errs.AppError)
	FindBy(id string) (*Customer, *errs.AppError)
}
