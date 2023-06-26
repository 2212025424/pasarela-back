package product

import (
	"github.com/google/uuid"

	"pasarelapago.go/model"
)

// ******************
// **** PORT OUT ****
// ******************

type Storage interface {
	All() (model.Products, error)
	ByID(ID uuid.UUID) (model.Product, error)
}

// *****************
// **** PORT IN ****
// *****************

type Product interface {
	All() (model.Products, error)
	ByID(ID uuid.UUID) (model.Product, error)
}
