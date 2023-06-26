package subscription

import "pasarelapago.go/model"

type Storage interface {
	Create(s *model.Subscription) error
	ByEmail(email string) (model.Subscriptions, error)
}

type Subscription interface {
	ByEmail(email string) (model.Subscriptions, error)
}
