package usecase

import (
	"context"
	"github.com/brendontj/didactic-eureka/core/repository"
	"github.com/brendontj/didactic-eureka/core/usecase/input"
	"github.com/google/uuid"
)

type DeleteCustomer struct {
	repo repository.Repository
}

func NewDeleteCustomer(repo repository.Repository) DeleteCustomer {
	return DeleteCustomer{
		repo: repo,
	}
}

func (uc *DeleteCustomer) Execute(ctx context.Context, data input.DeleteCustomerInput) error {
	id, err := uuid.Parse(data.ID)
	if err != nil {
		return err
	}
	err = uc.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
