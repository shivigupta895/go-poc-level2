package pubsub

import (
	"context"
	"encoding/json"
	"log"
	"payment-service/models"

	"cloud.google.com/go/pubsub"
)

var topicID = "tp-payment-events"

func PublishPaymentCreated(payment models.Payment) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, "potent-orbit-456909-f4")

	log.Printf("call line no. 18: %v", payment)
	if err != nil {
		log.Printf("PubSub error to publish payment event: %v", err)
		return
	}

	topic := client.Topic(topicID)
	data, _ := json.Marshal(payment)
	result := topic.Publish(ctx, &pubsub.Message{Data: data})

	_, err = result.Get(ctx)
	if err != nil {
		log.Printf("Failed to publish payment event: %v", err)
	}

	log.Printf("successfully call line no. 33: %v", payment)
}
