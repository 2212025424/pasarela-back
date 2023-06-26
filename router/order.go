package router

import (
	"database/sql"

	"github.com/labstack/echo/v4"

	"pasarelapago.go/domain/order"
	handler "pasarelapago.go/handler/order"
	storage "pasarelapago.go/storage/postgres/order"
)

func Order(e *echo.Echo, db *sql.DB) {
	store := storage.New(db)
	useCase := order.New(store)
	handler.NewRouter(e, useCase)
}
