package repository

import (
	"context"

	"github.com/fikrihkll/go-pubsub-publisher/infrastructure"
	"github.com/fikrihkll/go-pubsub-publisher/transport"
)

type OrderRepository struct {
	pubSub *infrastructure.PubSubService
}

func NewOrderRepository(pubSub *infrastructure.PubSubService) *OrderRepository {
	return &OrderRepository{
		pubSub: pubSub,
	}
}

func (r *OrderRepository) CreateOrder(c context.Context, data transport.Order) (id string, err error) {
	ctx := context.Background()

	id, err = r.pubSub.PublishMessage(ctx, r.pubSub.OrderPlacedTopic, data)
	if err != nil {
		return
	}

	return
}

func (r *OrderRepository) CancelOrder(c context.Context, data transport.Order) (id string, err error) {
	ctx := context.Background()

	id, err = r.pubSub.PublishMessage(ctx, r.pubSub.CancelOrderTopic, map[string]string{"well": "ok", "ok": "well"})

	if err != nil {
		return
	}

	return
}
