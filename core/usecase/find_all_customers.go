package usecase

import (
	"context"
	"github.com/brendontj/didactic-eureka/core/repository"
	"github.com/brendontj/didactic-eureka/core/usecase/common"
	"github.com/brendontj/didactic-eureka/core/usecase/output"
)

type FindAllCustomers struct {
	repo repository.Repository
}

func NewFindAllCustomers(repo repository.Repository) FindAllCustomers {
	return FindAllCustomers{
		repo: repo,
	}
}

func (uc *FindAllCustomers) Execute(ctx context.Context) (output.FindAllCustomersOutput, error) {
	customers, err := uc.repo.FindAll(ctx)
	if err != nil {
		return output.FindAllCustomersOutput{}, err
	}

	var customersOutput []output.CustomerOutput
	for _, customer := range customers {
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
		customersOutput = append(customersOutput, customerOutput)
	}

	return output.FindAllCustomersOutput{
		Customers: customersOutput,
	}, nil
}
