package domain

import (
	"abank/dto"
	"abank/errs"
)

type Account struct {
	Id          string  `db:"account_id"`
	CustomerId  string  `db:"customer_id"`
	OpeningDate string  `db:"opening_date"`
	AccountType string  `db:"account_type"`
	Amount      float64 `db:"amount"`
	Status      string  `db:"status:"`
}

func (a Account) ToNewAccountResponseDTO() dto.NewAccountResponse {
	return dto.NewAccountResponse{
		Id: a.Id,
	}
}

type AccountRepository interface {
	Create(acc Account) (*Account, *errs.AppError)
}
