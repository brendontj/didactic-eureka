package entity

import (
	"cloud.google.com/go/civil"
	"github.com/google/uuid"
	"time"
)

type Customer struct {
	ID        uuid.UUID
	Version   uuid.UUID
	Name      string
	Email     string
	Document  string
	Phone     string
	Address   Address
	BirthDate civil.Date
	CreatedAt time.Time
	UpdatedAt time.Time
}
