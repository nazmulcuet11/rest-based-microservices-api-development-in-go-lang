package app

import (
	"abank/dto"
	"abank/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type AccountHandler struct {
	service service.AccountService
}

func (h AccountHandler) createNewAccount(w http.ResponseWriter, r *http.Request) {
	var req dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	req.CustomerId = customerId

	res, appErr := h.service.NewAccount(req)
	if appErr != nil {
		writeResponse(w, appErr.Code, appErr.AsMessage())
		return
	}
	writeResponse(w, http.StatusCreated, res)
}
