package invoice

import "pasarelapago.go/model"

type Storage interface {
	Create(i *model.Invoice) error
	ByEmail(email string) (model.Invoices, error)
}

type Invoice interface {
	ByEmail(email string) (model.Invoices, error)
}
