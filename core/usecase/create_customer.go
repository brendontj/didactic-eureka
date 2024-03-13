package usecase

import (
	"context"
	"encoding/json"
	"github.com/brendontj/didactic-eureka/core/gateway"
	"github.com/brendontj/didactic-eureka/core/mapper"
	"github.com/brendontj/didactic-eureka/core/usecase/input"
	"github.com/brendontj/didactic-eureka/core/usecase/output"
)

type CreateCustomer struct {
	gateway   gateway.MessageGateway
	queueName string
}

func NewCreateCustomer(messageGateway gateway.MessageGateway, queueName string) CreateCustomer {
	return CreateCustomer{
		gateway:   messageGateway,
		queueName: queueName,
	}
}

func (uc *CreateCustomer) Execute(ctx context.Context, data input.CreateCustomerInput) (output.CreateCustomerOutput, error) {
	customerEntity, err := mapper.NewCustomerFromCreateCustomerInput(data)
	if err != nil {
		return output.CreateCustomerOutput{}, err
	}

	dataToSend, err := json.Marshal(customerEntity)
	if err != nil {
		return output.CreateCustomerOutput{}, err
	}

	if err = uc.gateway.Publish(ctx, uc.queueName, dataToSend); err != nil {
		return output.CreateCustomerOutput{}, err
	}

	return output.CreateCustomerOutput{
		ID:        customerEntity.ID.String(),
		Version:   customerEntity.Version.String(),
		CreatedAt: customerEntity.CreatedAt,
		UpdatedAt: customerEntity.UpdatedAt,
	}, nil
}
