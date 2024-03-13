package mapper

import (
	"cloud.google.com/go/civil"
	"github.com/brendontj/didactic-eureka/core/entity"
	"github.com/brendontj/didactic-eureka/core/usecase/input"
	"github.com/google/uuid"
	"time"
)

func NewCustomerFromCreateCustomerInput(data input.CreateCustomerInput) (entity.Customer, error) {
	birthDate, err := civil.ParseDate(data.BirthDate)
	if err != nil {
		return entity.Customer{}, err
	}

	return entity.Customer{
		ID:      uuid.New(),
		Version: uuid.New(),
		Name:    data.Name,
		Email:   data.Email,
		Phone:   data.Phone,
		Address: entity.Address{
			Street:       data.Address.Street,
			Number:       data.Address.Number,
			ZipCode:      data.Address.ZipCode,
			Complement:   data.Address.Complement,
			Neighborhood: data.Address.Neighborhood,
			City:         data.Address.City,
			State:        data.Address.State,
			Country:      data.Address.Country,
		},
		BirthDate: birthDate,
		Document:  data.Document,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func NewCustomerEntityUpdated(currentCustomerEntity entity.Customer, data input.UpdateCustomerInput) (entity.Customer, error) {
	birthDate, err := civil.ParseDate(data.BirthDate)
	if err != nil {
		return entity.Customer{}, err
	}

	return entity.Customer{
		ID:      currentCustomerEntity.ID,
		Version: uuid.New(),
		Name:    data.Name,
		Email:   data.Email,
		Phone:   data.Phone,
		Address: entity.Address{
			Street:       data.Address.Street,
			Number:       data.Address.Number,
			ZipCode:      data.Address.ZipCode,
			Complement:   data.Address.Complement,
			Neighborhood: data.Address.Neighborhood,
			City:         data.Address.City,
			State:        data.Address.State,
			Country:      data.Address.Country,
		},
		BirthDate: birthDate,
		Document:  data.Document,
		CreatedAt: currentCustomerEntity.CreatedAt,
		UpdatedAt: time.Now(),
	}, nil
}
