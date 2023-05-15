package app

import (
	"abank/domain"
	"abank/logger"
	"abank/service"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func Start() {
	sanityCheck()

	router := mux.NewRouter()

	// wiring
	db := createDBClient()
	ch := CustomerHandler{service: service.NewDefaultCustomerService(domain.NewCustomerRepositoryDB(db))}
	ah := AccountHandler{service: service.NewDefaultAccountService(domain.NewAccountRepositoryDB(db))}

	// define routes
	router.HandleFunc("/customers", ch.getCustomerList).Methods(http.MethodGet)
	router.HandleFunc("/customer/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customer/{customer_id:[0-9]+}/account", ah.createNewAccount).Methods(http.MethodPost)

	serverAddr := os.Getenv("SERVER_ADDR")
	serverPort := os.Getenv("SERVER_PORT")

	// start server
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%v:%v", serverAddr, serverPort), router))
	logger.Info(fmt.Sprintf("Application running at %v:%v", serverAddr, serverPort))
}

func sanityCheck() {
	variables := []string{
		"SERVER_ADDR",
		"SERVER_PORT",
		"DB_USER",
		"DB_PASS",
		"DB_ADDR",
		"DB_PORT",
		"DB_NAME",
	}

	for _, variable := range variables {
		if os.Getenv(variable) == "" {
			log.Fatal("Environment variable " + variable + " not defined!!")
		}
	}
}

func createDBClient() *sqlx.DB {
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
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	logger.Info(fmt.Sprintf("Connected to DB: %v", dbPath))

	return db
}
