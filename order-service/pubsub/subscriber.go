package pubsub

import (
	"context"
	"encoding/json"
	"log"
	"order-service/config"

	"cloud.google.com/go/pubsub"
)

func SubscribeToPaymentEvents() {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, config.GcpPojectId)
	if err != nil {
		log.Printf(`{"message":"Failed to create Pub/Sub client to subscribe payment event: %v", "service":"order", "severity":"ERROR"}`, err)
		return
	}

	log.Println(`{"message":"Client created in PubSub request to subscribe payment event", "service":"order", "severity":"INFO"}`)

	sub := client.Subscription(config.OrderSubId)
	err = sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		var event struct {
			OrderID string `json:"order_id"`
		}

		if err := json.Unmarshal(msg.Data, &event); err != nil {
			log.Printf(`{"message":"Error parsing message inside subscriber: %v", "service":"order", "severity":"ERROR"}`, err)
			msg.Nack()
			return
		}

		log.Printf(`{"message":"Successfully received payment event: %+v", "service":"order", "severity":"INFO"}`, event)
		config.UpdateOrder(event.OrderID)
		msg.Ack()
	})

	if err != nil {
		log.Fatalf(`{"message":"Error receiving messages inside subscriber: %v", "service":"order", "severity":"ERROR"}`, err)
	}
}
