package handler

import (
	"encoding/json"
	"github.com/brendontj/didactic-eureka/core/usecase"
	"github.com/brendontj/didactic-eureka/core/usecase/input"
	"net/http"
)

type FindCustomerByIdHandler struct {
	findCustomerByIDUseCase usecase.FindCustomerById
}

func NewFindCustomerByIdHandler(uc usecase.FindCustomerById) FindCustomerByIdHandler {
	return FindCustomerByIdHandler{
		findCustomerByIDUseCase: uc,
	}
}

func (h FindCustomerByIdHandler) Handle(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("id")
	if userID == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	outputData, err := h.findCustomerByIDUseCase.Execute(r.Context(), input.FindCustomerByIdInput{ID: userID})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(outputData)
}
