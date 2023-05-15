package domain

import (
	"abank/dto"
	"abank/errs"
)

type Account struct {
	Id          string  `db:"account_id"`
	CustomerId  string  `db:"customer_id"`
	OpeningDate string  `db:"opening_date"`
	Type        string  `db:"account_type"`
	Amount      float64 `db:"amount"`
	Status      string  `db:"status"`
}

func (a Account) ToNewAccountResponseDTO() dto.NewAccountResponse {
	return dto.NewAccountResponse{
		Id: a.Id,
	}
}

func (a Account) CanWithDrawAmount(amount float64) bool {
	return amount <= a.Amount
}

type AccountRepository interface {
	Create(acc Account) (*Account, *errs.AppError)
	FindBy(id string) (*Account, *errs.AppError)
	SaveTransaction(t Transaction) (*Transaction, *errs.AppError)
}
