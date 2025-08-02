package pubsub

import (
	"context"
	"encoding/json"
	"log"
	"payment-service/config"
	"payment-service/models"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SubscribeToOrderEvents(db *gorm.DB) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, config.GcpPojectId)
	if err != nil {
		log.Printf(`{"message":"Failed to create Pub/Sub client to subscribe order event: %v", "service":"payment", "severity":"ERROR"}`, err)
	}

	log.Println(`{"message":"Client created in PubSub request to subscribe order event", "service":"payment", "severity":"INFO"}`)

	sub := client.Subscription(config.PaymentSubId)
	err = sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		var event struct {
			ID string `json:"id"`
		}

		if err := json.Unmarshal(msg.Data, &event); err != nil {
			log.Printf(`{"message":"Error parsing message inside subscriber: %v", "service":"payment", "severity":"ERROR"}`, err)
			msg.Nack()
			return
		}

		log.Printf(`{"message":"Successfully received order event: %+v", "service":"payment", "severity":"INFO"}`, event)

		payment := models.Payment{
			ID:        uuid.New().String(),
			OrderID:   event.ID,
			Status:    "PAID", // Simulate payment success
			UpdatedAt: time.Now(),
		}

		if err := db.Create(&payment).Error; err != nil {
			log.Printf(`{"message":"Failed to create payment: %v", "service":"payment", "severity":"ERROR"}`, err)
			msg.Nack()
			return
		}

		log.Printf(`{"message":"Payment created successfully: %v", "service":"payment", "severity":"INFO"}`, payment)

		PublishPaymentCreated(payment)

		msg.Ack()
	})

	if err != nil {
		log.Fatalf(`{"message":"Error receiving messages inside subscriber: %v", "service":"payment", "severity":"ERROR"}`, err)
	}
}
