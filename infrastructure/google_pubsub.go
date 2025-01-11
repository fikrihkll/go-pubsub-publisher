package infrastructure

import (
	"context"
	"encoding/json"
	"fmt"

	"cloud.google.com/go/pubsub"
)

type PubSubService struct {
	Client             *pubsub.Client
	OrderPlacedTopic   *pubsub.Topic
	CancelOrderTopic *pubsub.Topic
}

func NewPubSub(ctx context.Context) (*PubSubService, error){
	client, err := pubsub.NewClient(ctx, "golang-practice-447504")
	if err != nil {
		return nil, fmt.Errorf("failed to create Pub/Sub client: %w", err)
	}

	service := &PubSubService{
		Client: client,
		OrderPlacedTopic: client.Topic("order-placed-topic"),
		CancelOrderTopic: client.Topic("cancel-order-topic"),
	}
	return service, nil
}

func (p* PubSubService) PublishMessage(ctx context.Context, topic *pubsub.Topic, data interface{}) (string, error) {
	messageBytes, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("failed to marshal message: %v", err)
	}

	result := topic.Publish(ctx, &pubsub.Message{
		Data: messageBytes,
	})
	id, err := result.Get(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to publish message: %v", err)
	}

	return id, nil
}
