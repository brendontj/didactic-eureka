package handler

import (
	"github.com/brendontj/didactic-eureka/core/usecase"
	"github.com/brendontj/didactic-eureka/core/usecase/input"
	"net/http"
)

type DeleteCustomerHandler struct {
	DeleteCustomerUseCase usecase.DeleteCustomer
}

func NewDeleteCustomerHandler(uc usecase.DeleteCustomer) DeleteCustomerHandler {
	return DeleteCustomerHandler{
		DeleteCustomerUseCase: uc,
	}
}

func (h DeleteCustomerHandler) Handle(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("id")
	if userID == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	if err := h.DeleteCustomerUseCase.Execute(r.Context(), input.DeleteCustomerInput{ID: userID}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	w.WriteHeader(http.StatusOK)
}
