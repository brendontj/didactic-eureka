package usecase

import (
	"context"
	"fmt"
	"github.com/brendontj/didactic-eureka/core/mapper"
	"github.com/brendontj/didactic-eureka/core/repository"
	"github.com/brendontj/didactic-eureka/core/usecase/input"
	"github.com/brendontj/didactic-eureka/core/usecase/output"
	"github.com/google/uuid"
)

var (
	ErrVersionMismatch = fmt.Errorf("version mismatch")
)

type UpdateCustomer struct {
	repo repository.Repository
}

func NewUpdateCustomer(repo repository.Repository) UpdateCustomer {
	return UpdateCustomer{
		repo: repo,
	}
}

func (uc *UpdateCustomer) Execute(ctx context.Context, data input.UpdateCustomerInput) (output.UpdateCustomerOutput, error) {
	id, err := uuid.Parse(data.ID)
	if err != nil {
		return output.UpdateCustomerOutput{}, err
	}

	currentCustomerEntity, err := uc.repo.FindById(ctx, id)
	if err != nil {
		return output.UpdateCustomerOutput{}, err
	}

	if currentCustomerEntity.Version.String() != data.Version {
		return output.UpdateCustomerOutput{}, ErrVersionMismatch
	}

	updatedEntity, err := mapper.NewCustomerEntityUpdated(currentCustomerEntity, data)
	if err != nil {
		return output.UpdateCustomerOutput{}, err
	}

	if err = uc.repo.Update(ctx, updatedEntity, currentCustomerEntity.Version); err != nil {
		return output.UpdateCustomerOutput{}, err
	}

	return output.UpdateCustomerOutput{
		ID:        updatedEntity.ID.String(),
		Version:   updatedEntity.Version.String(),
		CreatedAt: updatedEntity.CreatedAt,
		UpdatedAt: updatedEntity.UpdatedAt,
	}, nil
}
