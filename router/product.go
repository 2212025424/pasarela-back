package router

import (
	"database/sql"

	"github.com/labstack/echo/v4"

	"pasarelapago.go/domain/product"
	handler "pasarelapago.go/handler/product"
	storage "pasarelapago.go/storage/postgres/product"
)

func Product(e *echo.Echo, db *sql.DB) {
	store := storage.New(db)
	useCase := product.New(store)
	handler.NewRouter(e, useCase)
}
