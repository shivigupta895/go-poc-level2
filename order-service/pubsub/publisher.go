package pubsub

import (
	"context"
	"encoding/json"
	"log"
	"order-service/config"
	"order-service/models"

	"cloud.google.com/go/pubsub"
)

func PublishOrderCreated(order models.Order) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, config.GcpPojectId)
	if err != nil {
		log.Printf(`{"message":"PubSub client error to publish order event: %v", "service":"order", "severity":"ERROR"}`, err)
		return
	}

	log.Println(`{"message":"Client created in PubSub request", "service":"order", "severity":"INFO"}`)

	topic := client.Topic(config.OrderTopicId)
	data, _ := json.Marshal(order)
	result := topic.Publish(ctx, &pubsub.Message{Data: data})

	_, err = result.Get(ctx)
	if err != nil {
		log.Printf(`{"message":"Failed to publish order event: %v", "service":"order", "severity":"ERROR"}`, err)
	}

	log.Printf(`{"message":"Successfully published order event: %v", "service":"order", "severity":"INFO"}`, order)
}
