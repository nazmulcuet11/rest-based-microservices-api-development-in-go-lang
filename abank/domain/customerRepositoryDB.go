package domain

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"abank/errs"
	"abank/logger"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDB struct {
	db *sqlx.DB
}

func NewCustomerRepositoryDB() CustomerRepositoryDB {
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	dbCred := fmt.Sprintf("%v:%v", dbUser, dbPass)
	dbPath := fmt.Sprintf("tcp(%v:%v)/%v", dbAddr, dbPort, dbName)
	connString := fmt.Sprintf("%v@%v", dbCred, dbPath)
	db, err := sqlx.Open("mysql", connString)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	logger.Info(fmt.Sprintf("Connected to DB: %v", dbPath))
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
