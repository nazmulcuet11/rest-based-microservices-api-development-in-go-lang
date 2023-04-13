package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	// define routes
	router.HandleFunc("/api/time", getCurrentTime).Methods(http.MethodGet)

	// start server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
