package dto

import (
	"abank/errs"
	"fmt"
)

const (
	WITHDRAWL = "withdrawl"
	DEPOSIT   = "deposit"
)

type TransactionRequest struct {
	CustomerId      string
	AccountId       string
	TransactionType string  `json:"transaction_type"`
	Amount          float64 `json:"amount"`
}

func (r TransactionRequest) Validate() *errs.AppError {

	if r.TransactionType != WITHDRAWL && r.TransactionType != DEPOSIT {
		errMessage := fmt.Sprintf("Account type can be either %s or %s", WITHDRAWL, DEPOSIT)
		return errs.NewValidationError(errMessage)
	}

	if r.Amount <= 0 {
		return errs.NewValidationError("Ammount must be greater than zero")
	}

	return nil
}

func (r TransactionRequest) IsWithdrawl() bool {
	return r.TransactionType == WITHDRAWL
}
