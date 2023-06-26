package paypal

import (
	"net/http"

	"github.com/google/uuid"

	"pasarelapago.go/model"
)

// PayPal is implemented by the Handler
type PayPal interface {
	ProcessRequest(header http.Header, body []byte) error
}

// PortsOut

type Order interface {
	ByID(ID uuid.UUID) (model.Order, error)
}

type Subscription interface {
	Create(s *model.Subscription) error
}

type Invoice interface {
	Create(order *model.Order, subscriptionID uuid.UUID) error
}
