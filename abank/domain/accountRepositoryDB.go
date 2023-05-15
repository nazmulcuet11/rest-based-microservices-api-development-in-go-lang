package domain

import (
	"abank/errs"
	"abank/logger"
	"database/sql"
	"log"
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

	result, err := r.db.Exec(
		queryString,
		acc.CustomerId,
		acc.OpeningDate,
		acc.Type,
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

func (r AccountRepositoryDB) FindBy(id string) (*Account, *errs.AppError) {
	queryString := `SELECT
						account_id,
						customer_id,
						opening_date,
						account_type,
						amount, 
						status
					FROM accounts
					WHERE account_id = ?`

	var account Account
	err := r.db.Get(&account, queryString, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NotFoundError("Account not found")
		}

		log.Println("Error converting row to account: ", err)
		return nil, errs.InternalServerError("Unexpected databse error")
	}

	return &account, nil
}

func (r AccountRepositoryDB) SaveTransaction(t Transaction) (*Transaction, *errs.AppError) {
	tx, err := r.db.Begin()
	if err != nil {
		logger.Error("Error while starting transaction: " + err.Error())
		return nil, errs.InternalServerError("Unexpected databse error")
	}

	queryString := `
	INSERT INTO transactions (
		account_id,
		amount, 
		transaction_type,
		transaction_date
	) VALUES (
		?,
		?,
		?,
		?
	)`

	result, err := tx.Exec(queryString, t.AccountId, t.Amount, t.Type, t.Date)
	if err != nil {
		tx.Rollback()
		logger.Error("Error while saving transaction: " + err.Error())
		return nil, errs.InternalServerError("Unexpected databse error")
	}

	if t.IsWithdrawl() {
		queryString = `UPDATE accounts SET amount = amount - ? WHERE account_id = ?`
		_, err = tx.Exec(queryString, t.Amount, t.AccountId)
	} else {
		queryString = `UPDATE accounts SET amount = amount + ? WHERE account_id = ?`
		_, err = tx.Exec(queryString, t.Amount, t.AccountId)
	}

	if err != nil {
		tx.Rollback()
		logger.Error("Error while updating amount: " + err.Error())
		return nil, errs.InternalServerError("Unexpected databse error")
	}

	err = tx.Commit()
	if err != nil {
		logger.Error("Error while commiting transaction: " + err.Error())
		return nil, errs.InternalServerError("Unexpected databse error")
	}

	transactionId, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting las transaction Id: " + err.Error())
		return nil, errs.InternalServerError("Unexpected databse error")
	}
	account, appErr := r.FindBy(t.AccountId)
	if appErr != nil {
		return nil, appErr
	}

	t.Id = strconv.FormatInt(transactionId, 10)
	t.Amount = account.Amount

	return &t, nil
}
