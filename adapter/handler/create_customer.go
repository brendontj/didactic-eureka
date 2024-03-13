package handler

import (
	"encoding/json"
	"github.com/brendontj/didactic-eureka/core/usecase"
	"github.com/brendontj/didactic-eureka/core/usecase/input"
	"net/http"
)

type CreateCustomerHandler struct {
	CreateCustomerUseCase usecase.CreateCustomer
}

func NewCreateCustomerHandler(uc usecase.CreateCustomer) CreateCustomerHandler {
	return CreateCustomerHandler{
		CreateCustomerUseCase: uc,
	}
}

func (h CreateCustomerHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var inputData input.CreateCustomerInput

	err := json.NewDecoder(r.Body).Decode(&inputData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	outputData, err := h.CreateCustomerUseCase.Execute(r.Context(), inputData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(outputData)
}
