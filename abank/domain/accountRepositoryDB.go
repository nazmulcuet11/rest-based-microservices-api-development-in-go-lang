package domain

import (
	"abank/errs"
	"abank/logger"
	"fmt"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDB struct {
	db *sqlx.DB
}

func NewAccountRepositoryDB(db *sqlx.DB) AccountRepositoryDB {
	return AccountRepositoryDB{db: db}
}

func (r AccountRepositoryDB) Create(acc Account) (*Account, *errs.AppError) {
	queryString := `
	INSERT INTO accounts (
		customer_id,
		opening_date,
		account_type,
		amount, 
		status
	) VALUES (
		?,
		?,
		?,
		?,
		?
	)`

	fmt.Println(queryString)
	result, err := r.db.Exec(
		queryString,
		acc.CustomerId,
		acc.OpeningDate,
		acc.AccountType,
		acc.Amount,
		acc.Status,
	)

	if err != nil {
		logger.Error("Error creating new account: " + err.Error())
		return nil, errs.InternalServerError("Unexpected error from database")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error getting last inserted id: " + err.Error())
		return nil, errs.InternalServerError("Unexpected error from database")
	}

	acc.Id = strconv.FormatInt(id, 10)
	return &acc, nil
}
