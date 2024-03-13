package input

import "github.com/brendontj/didactic-eureka/core/usecase/common"

type CreateCustomerInput struct {
	*common.CustomerData
}

type DeleteCustomerInput struct {
	ID string
}

type FindCustomerByIdInput struct {
	ID string
}

type UpdateCustomerInput struct {
	ID      string `json:"id"`
	Version string `json:"version"`
	*common.CustomerData
}
