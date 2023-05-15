package app

import (
	"abank/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type CustomerHandler struct {
	service service.CustomerService
}

func NewCustomerHandler(service service.CustomerService) CustomerHandler {
	return CustomerHandler{service: service}
}

func (h *CustomerHandler) getCustomerList(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	customers, err := h.service.GetAllCustomers(status)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customers)
	}
}

func (h *CustomerHandler) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	customer, err := h.service.GetCustmerBy(customerId)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customer)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		panic(err)
	}
}
