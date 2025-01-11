package main

import (
	"context"
	"log"
	"net/http"

	handler "github.com/fikrihkll/go-pubsub-publisher/http"
	"github.com/fikrihkll/go-pubsub-publisher/infrastructure"
	"github.com/fikrihkll/go-pubsub-publisher/repository"
	"github.com/labstack/echo/v4"
)

func initPubSub(ctx context.Context) (service *infrastructure.PubSubService) {
	service, err := infrastructure.NewPubSub(ctx)
	if err != nil {
		log.Fatalf("Failed to create Pub/Sub client: %v", err)
	}
	return
}

func main() {
	ctx := context.Background()

	e := echo.New()

	pubSubService := initPubSub(ctx)
	orderRepository := repository.NewOrderRepository(pubSubService)
	app := handler.NewOrderApplication(orderRepository)

	defer pubSubService.Client.Close()

	e.GET("/publish/status", func(c echo.Context) error { return c.JSON(http.StatusOK, map[string]string{"status": "OK!"}) })
	e.POST("/publish/order", app.CreateOrder)
	e.POST("/publish/cancel", app.CancelOrder)

	// Start the server
	log.Println("Publisher server is running on http://localhost:8080")
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
