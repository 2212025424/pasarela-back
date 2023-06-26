package product

import (
	"github.com/google/uuid"

	"pasarelapago.go/model"
)

type UseCase struct {
	storage Storage
}

func New(s Storage) UseCase {
	return UseCase{storage: s}
}

func (uc UseCase) All() (model.Products, error) {
	return uc.storage.All()
}

func (uc UseCase) ByID(ID uuid.UUID) (model.Product, error) {
	return uc.storage.ByID(ID)
}
