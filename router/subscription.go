package router

import (
	"database/sql"

	"github.com/labstack/echo/v4"

	"pasarelapago.go/domain/subscription"
	handler "pasarelapago.go/handler/subscription"
	storage "pasarelapago.go/storage/postgres/subscription"
)

func Subscription(e *echo.Echo, db *sql.DB) {
	store := storage.New(db)
	useCase := subscription.New(store)
	handler.NewRouter(e, useCase)
}
