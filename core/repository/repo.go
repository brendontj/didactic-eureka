package repository

import (
	"context"
	"github.com/brendontj/didactic-eureka/core/entity"
	"github.com/google/uuid"
)

type Repository interface {
	FindAll(ctx context.Context) ([]entity.Customer, error)
	FindById(ctx context.Context, id uuid.UUID) (entity.Customer, error)
	Save(ctx context.Context, customer entity.Customer) error
	Update(ctx context.Context, customer entity.Customer, currentCustomerVersion uuid.UUID) error
	Delete(ctx context.Context, id uuid.UUID) error
}
