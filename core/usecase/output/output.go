package output

import (
	"github.com/brendontj/didactic-eureka/core/usecase/common"
	"time"
)

type CreateCustomerOutput struct {
	ID        string    `json:"id"`
	Version   string    `json:"version"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type FindAllCustomersOutput struct {
	Customers []CustomerOutput `json:"customers"`
}

type CustomerOutput struct {
	ID      string `json:"id"`
	Version string `json:"version"`
	*common.CustomerData
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type FindCustomerByIdOutput struct {
	*CustomerOutput
}

type UpdateCustomerOutput struct {
	ID        string    `json:"id"`
	Version   string    `json:"version"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
