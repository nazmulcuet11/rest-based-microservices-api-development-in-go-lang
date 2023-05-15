package dto

import (
	"abank/errs"
)

type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Ammount     float64 `json:"ammount"`
}

func (r NewAccountRequest) Validate() *errs.AppError {
	if r.Ammount < 5000 {
		return errs.NewValidationError("Need to deposit at least 5000")
	}

	if r.AccountType == "" {
		return errs.NewValidationError("missing requried field `account_type`")
	}

	if r.AccountType != "saving" && r.AccountType != "checking" {
		return errs.NewValidationError("Account type can be either saving or checking")
	}

	return nil
}
