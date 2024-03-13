package handler

import (
	"encoding/json"
	"github.com/brendontj/didactic-eureka/core/usecase"
	"net/http"
)

type FindAllCustomersHandler struct {
	findAllCustomersUseCase usecase.FindAllCustomers
}

func NewFindAllCustomersHandler(uc usecase.FindAllCustomers) FindAllCustomersHandler {
	return FindAllCustomersHandler{
		findAllCustomersUseCase: uc,
	}
}

func (h FindAllCustomersHandler) Handle(w http.ResponseWriter, r *http.Request) {
	outputData, err := h.findAllCustomersUseCase.Execute(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(outputData)
}
