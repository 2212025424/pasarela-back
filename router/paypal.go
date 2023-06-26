package router

import (
	"database/sql"

	"github.com/labstack/echo/v4"

	"pasarelapago.go/domain/invoice"
	"pasarelapago.go/domain/order"
	"pasarelapago.go/domain/paypal"
	"pasarelapago.go/domain/subscription"

	storageInvoice "pasarelapago.go/storage/postgres/invoice"
	storageOrder "pasarelapago.go/storage/postgres/order"
	storageSubscription "pasarelapago.go/storage/postgres/subscription"

	handler "pasarelapago.go/handler/paypal"
)

func PayPal(e *echo.Echo, db *sql.DB) {
	useCaseOrder := buildOrder(db)
	useCaseSubs := buildSubs(db)
	useCaseInvoice := buildInvoice(db)
	useCasePayPal := paypal.New(useCaseOrder, useCaseSubs, useCaseInvoice)

	handler.NewRouter(e, useCasePayPal)
}

func buildOrder(db *sql.DB) paypal.Order {
	store := storageOrder.New(db)
	return order.New(store)
}

func buildSubs(db *sql.DB) paypal.Subscription {
	store := storageSubscription.New(db)
	return subscription.New(store)
}

func buildInvoice(db *sql.DB) paypal.Invoice {
	store := storageInvoice.New(db)
	return invoice.New(store)
}
