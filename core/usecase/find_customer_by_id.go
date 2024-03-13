package usecase

import (
	"context"
	"github.com/brendontj/didactic-eureka/core/repository"
	"github.com/brendontj/didactic-eureka/core/usecase/common"
	"github.com/brendontj/didactic-eureka/core/usecase/input"
	"github.com/brendontj/didactic-eureka/core/usecase/output"
	"github.com/google/uuid"
)

type FindCustomerById struct {
	repo repository.Repository
}

func NewFindCustomerById(repo repository.Repository) FindCustomerById {
	return FindCustomerById{
		repo: repo,
	}
}

func (uc *FindCustomerById) Execute(ctx context.Context, input input.FindCustomerByIdInput) (output.FindCustomerByIdOutput, error) {
	id, err := uuid.Parse(input.ID)
	if err != nil {
		return output.FindCustomerByIdOutput{}, err
	}

	customer, err := uc.repo.FindById(ctx, id)
	if err != nil {
		return output.FindCustomerByIdOutput{}, err
	}

	customerOutput := output.CustomerOutput{
		ID:      customer.ID.String(),
		Version: customer.Version.String(),
		CustomerData: &common.CustomerData{
			Name:  customer.Name,
			Email: customer.Email,
			Phone: customer.Phone,
			Address: common.CustomerAddressData{
				Street:       customer.Address.Street,
				Number:       customer.Address.Number,
				ZipCode:      customer.Address.ZipCode,
				Neighborhood: customer.Address.Neighborhood,
				City:         customer.Address.City,
				State:        customer.Address.State,
				Country:      customer.Address.Country,
				Complement:   customer.Address.Complement,
			},
			BirthDate: customer.BirthDate.String(),
			Document:  customer.Document,
		},
		CreatedAt: customer.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: customer.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	return output.FindCustomerByIdOutput{
		CustomerOutput: &customerOutput,
	}, nil
}
