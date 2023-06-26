package order

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

func (uc UseCase) Create(o *model.Order) error {
	o.ID = uuid.New()
	return uc.storage.Create(o)
}

func (uc UseCase) ByID(ID uuid.UUID) (model.Order, error) {
	return uc.storage.ByID(ID)
}
