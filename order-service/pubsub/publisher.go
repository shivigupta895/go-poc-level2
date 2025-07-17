package pubsub

import (
	"context"
	"encoding/json"
	"log"
	"order-service/models"

	"cloud.google.com/go/pubsub"
)

var topicID = "tp-order-events"

func PublishOrderCreated(order models.Order) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, "potent-orbit-456909-f4")
	if err != nil {
		log.Printf("PubSub client error to publish order event: %v", err)
		return
	}

	topic := client.Topic(topicID)
	data, _ := json.Marshal(order)
	log.Printf("Order created data line 24 : %v", order)
	result := topic.Publish(ctx, &pubsub.Message{Data: data})

	_, err = result.Get(ctx)
	if err != nil {
		log.Printf("Failed to publish order event: %v", err)
	}
}
