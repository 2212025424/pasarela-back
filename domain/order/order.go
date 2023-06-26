package order

import (
	"github.com/google/uuid"

	"pasarelapago.go/model"
)

type Storage interface {
	Create(o *model.Order) error
	ByID(ID uuid.UUID) (model.Order, error)
}

type Order interface {
	Create(o *model.Order) error
	ByID(ID uuid.UUID) (model.Order, error)
}
