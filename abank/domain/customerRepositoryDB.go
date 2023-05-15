package domain

import (
	"database/sql"
	"log"

	"abank/errs"
	"abank/logger"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDB struct {
	db *sqlx.DB
}

func NewCustomerRepositoryDB(db *sqlx.DB) CustomerRepositoryDB {
	return CustomerRepositoryDB{db: db}
}

func (r CustomerRepositoryDB) FindAll(status string) ([]Customer, *errs.AppError) {
	queryString := `SELECT
						customer_id,
						first_name, 
						last_name, 
						date_of_birth,
						city, 
	  					zipcode,
	   					status 
					FROM customers`

	customers := make([]Customer, 0)
	var err error
	if status == "" {
		err = r.db.Select(&customers, queryString)
	} else {
		queryString = queryString + " WHERE status = ?"
		err = r.db.Select(&customers, queryString, status)
	}

	if err != nil {
		logger.Error("Error getting all customers: " + err.Error())
		return nil, errs.InternalServerError("Unexpected databse error")
	}
	return customers, nil
}

func (r CustomerRepositoryDB) FindBy(id string) (*Customer, *errs.AppError) {
	queryString := `SELECT
						customer_id,
						first_name, 
						last_name, 
						date_of_birth,
						city, 
						zipcode,
					status 
					FROM customers
					WHERE customer_id = ?`

	var customer Customer
	err := r.db.Get(&customer, queryString, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NotFoundError("Customer not found")
		}

		log.Println("Error converting row to customer: ", err)
		return nil, errs.InternalServerError("Unexpected databse error")
	}

	return &customer, nil
}
