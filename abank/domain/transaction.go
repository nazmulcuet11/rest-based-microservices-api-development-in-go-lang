package domain

import "abank/dto"

type TransactionType string

const (
	WITHDRAWL TransactionType = "withdrawl"
	DEPOSIT   TransactionType = "deposit"
)

type Transaction struct {
	Id        string          `db:"transaction_id"`
	AccountId string          `db:"account_id"`
	Amount    float64         `db:"amount"`
	Type      TransactionType `db:"transaction_type"`
	Date      string          `db:"transaction_date"`
}

func (t Transaction) IsWithdrawl() bool {
	return t.Type == WITHDRAWL
}

func (t Transaction) ToTransactionResponseDTO() dto.TransactionResponse {
	return dto.TransactionResponse{
		TransactionId:   t.Id,
		AccountId:       t.AccountId,
		NewBalance:      t.Amount,
		TransactionType: string(t.Type),
		TransactionDate: t.Date,
	}
}
