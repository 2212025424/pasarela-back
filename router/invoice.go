package router

import (
	"database/sql"

	"github.com/labstack/echo/v4"

	"pasarelapago.go/domain/invoice"
	handler "pasarelapago.go/handler/invoice"
	storage "pasarelapago.go/storage/postgres/invoice"
)

func Invoice(e *echo.Echo, db *sql.DB) {
	store := storage.New(db)
	useCase := invoice.New(store)
	handler.NewRouter(e, useCase)
}
