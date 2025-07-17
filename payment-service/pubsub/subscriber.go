package pubsub

import (
	"context"
	"encoding/json"
	"log"
	"payment-service/models"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderEvent struct {
	ID     string  `json:"id"`
	Item   string  `json:"item"`
	Amount float64 `json:"amount"`
	Status string  `json:"status"`
}

func SubscribeToOrderEvents(db *gorm.DB) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, "potent-orbit-456909-f4")
	if err != nil {
		log.Fatalf("Failed to create Pub/Sub client: %v", err)
	}

	sub := client.Subscription("sb-payment-events")
	err = sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		var event OrderEvent
		if err := json.Unmarshal(msg.Data, &event); err != nil {
			log.Printf("Error parsing message: %v", err)
			msg.Nack()
			return
		}

		log.Printf("Received order event: %+v", event)

		payment := models.Payment{
			ID:        uuid.New().String(),
			OrderID:   event.ID,
			Status:    "PAID", // Simulate payment success
			UpdatedAt: time.Now(),
		}

		if err := db.Create(&payment).Error; err != nil {
			log.Printf("Failed to save payment: %v", err)
			msg.Nack()
			return
		}

		log.Printf("Get payment here: %v", payment)

		PublishPaymentCreated(payment)

		msg.Ack()
	})

	if err != nil {
		log.Fatalf("Error receiving messages: %v", err)
	}
}
