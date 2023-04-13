package domain

import "abank/errs"

type Customer struct {
	Id          string `db:"customer_id"`
	FirstName   string `db:"first_name"`
	LastName    string `db:"last_name"`
	City        string `db:"city"`
	Zipcode     string `db:"zipcode"`
	DateOfBirth string `db:"date_of_birth"`
	Status      string `db:"status"`
}

type CustomerRepository interface {
	FindAll(status string) ([]Customer, *errs.AppError)
	FindBy(id string) (*Customer, *errs.AppError)
}
