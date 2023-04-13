package app

import (
	"abank/domain"
	"abank/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()
	// wiring
	ch := CustomerHandler{service: service.NewDefaultCustomerService(domain.NewCustomerRepositoryDB())}

	// define routes
	router.HandleFunc("/", greet).Methods(http.MethodGet)
	router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/customers", ch.getCustomerList).Methods(http.MethodGet)
	router.HandleFunc("/customer/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	// start server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
