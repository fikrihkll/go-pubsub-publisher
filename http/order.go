package http

import (
	"fmt"
	"net/http"

	"github.com/fikrihkll/go-pubsub-publisher/repository"
	"github.com/fikrihkll/go-pubsub-publisher/transport"
	"github.com/labstack/echo/v4"
)

type OrderApplication struct {
	orderRepository *repository.OrderRepository
}

func NewOrderApplication(orderRepository *repository.OrderRepository) *OrderApplication {
	return &OrderApplication{
		orderRepository: orderRepository,
	}
}

func (app OrderApplication) CreateOrder(c echo.Context) error {
	var order transport.Order
	if err := c.Bind(&order); err != nil {
		fmt.Printf("pAYLOAD: %s | %s", order.OrderID, order.UserEmail)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request payload"})
	}

	id, err := app.orderRepository.CreateOrder(c.Request().Context(), order)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "something went wrong!"})
	}

	return c.JSON(http.StatusOK, map[string]string{"messageID": id})
}

func (app OrderApplication) CancelOrder(c echo.Context) error {
	var order transport.Order
	if err := c.Bind(&order); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request payload"})
	}

	id, err := app.orderRepository.CancelOrder(c.Request().Context(), order)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "something went wrong!"})
	}

	return c.JSON(http.StatusOK, map[string]string{"messageID": id})
}
