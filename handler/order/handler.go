package order

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"pasarelapago.go/domain/order"
	"pasarelapago.go/model"
)

type Handler struct {
	useCase order.Order
}

func New(uc order.Order) Handler {
	return Handler{useCase: uc}
}

func (h Handler) Create(c echo.Context) error {
	o := model.Order{}
	err := c.Bind(&o)
	if err != nil {
		msg := map[string]string{
			"error":    "la estructura de la orden no es correcta",
			"internal": err.Error(),
		}

		return c.JSON(http.StatusBadRequest, msg)
	}

	err = h.useCase.Create(&o)
	if err != nil {
		msg := map[string]string{
			"error":    "no pudimos crear la orden",
			"internal": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, msg)
	}

	return c.JSON(http.StatusOK, map[string]model.Order{"message": o})
}

func (h Handler) ByID(c echo.Context) error {
	ID := c.Param("id")
	data, err := h.useCase.ByID(uuid.MustParse(ID))
	if err != nil {
		msg := map[string]string{
			"error":    "no pudimos consultar la orden",
			"internal": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, msg)
	}

	msg := map[string]model.Order{"data": data}
	return c.JSON(http.StatusOK, msg)
}
