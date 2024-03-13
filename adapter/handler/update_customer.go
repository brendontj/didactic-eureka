package handler

import (
	"encoding/json"
	"github.com/brendontj/didactic-eureka/core/usecase"
	"github.com/brendontj/didactic-eureka/core/usecase/common"
	"github.com/brendontj/didactic-eureka/core/usecase/input"
	"net/http"
)

type UpdateCustomerHandler struct {
	UpdateCustomerUseCase usecase.UpdateCustomer
}

func NewUpdateCustomerHandler(uc usecase.UpdateCustomer) UpdateCustomerHandler {
	return UpdateCustomerHandler{
		UpdateCustomerUseCase: uc,
	}
}

func (h UpdateCustomerHandler) Handle(w http.ResponseWriter, r *http.Request) {
	customerID := r.Header.Get("id")
	customerVersion := r.Header.Get("version")

	if customerID == "" || customerVersion == "" {
		http.Error(w, "id and version are required", http.StatusBadRequest)
		return
	}

	var inputData input.CreateCustomerInput
	err := json.NewDecoder(r.Body).Decode(&inputData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updateCustomerOutput, err := h.UpdateCustomerUseCase.Execute(r.Context(), input.UpdateCustomerInput{
		ID:      customerID,
		Version: customerVersion,
		CustomerData: &common.CustomerData{
			Name:      inputData.Name,
			Email:     inputData.Email,
			Phone:     inputData.Phone,
			Address:   inputData.Address,
			BirthDate: inputData.BirthDate,
			Document:  inputData.Document,
		},
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(updateCustomerOutput)
}
