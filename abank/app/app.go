package app

import (
	"abank/domain"
	"abank/logger"
	"abank/service"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func Start() {
	sanityCheck()

	router := mux.NewRouter()
	// wiring
	ch := CustomerHandler{service: service.NewDefaultCustomerService(domain.NewCustomerRepositoryDB())}

	// define routes
	router.HandleFunc("/customers", ch.getCustomerList).Methods(http.MethodGet)
	router.HandleFunc("/customer/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	serverAddr := os.Getenv("SERVER_ADDR")
	serverPort := os.Getenv("SERVER_PORT")

	// start server
	logger.Info(fmt.Sprintf("Application running at %v:%v", serverAddr, serverPort))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%v:%v", serverAddr, serverPort), router))
}

func sanityCheck() {
	variables := []string{
		// server config
		"SERVER_ADDR",
		"SERVER_PORT",
		// db config
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
