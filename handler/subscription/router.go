package subscription

import (
	"github.com/labstack/echo/v4"

	"pasarelapago.go/domain/subscription"
)

const (
	path        = "/v1/subscriptions"
	pathByEmail = "/:email"
)

func NewRouter(e *echo.Echo, useCase subscription.Subscription) {
	handler := New(useCase)

	g := e.Group(path)
	g.GET(pathByEmail, handler.ByEmail)
}
