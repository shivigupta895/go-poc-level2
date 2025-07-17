package pubsub

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"order-service/models"

	"cloud.google.com/go/pubsub"
	"gorm.io/gorm"
)

type PaymentEvent struct {
	OrderID string `json:"order_id"`
	Status  string `json:"status"`
}

func SubscribeToPaymentEvents(db *gorm.DB) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, "potent-orbit-456909-f4")
	if err != nil {
		log.Fatalf("Failed to create Pub/Sub client: %v", err)
	}

	sub := client.Subscription("sb-order-events")

	err = sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		var event PaymentEvent
		if err := json.Unmarshal(msg.Data, &event); err != nil {
			log.Printf("Error parsing message: %v", err)
			msg.Nack()
			return
		}

		log.Printf("Received payment event: %+v", event)

		order := models.Order{
			ID:     event.OrderID,
			Status: "PAID", // Simulate payment success
		}

		// callPatchEndpoint("http://localhost:8080/orders/"+event.OrderID+"/status", &order)
		callPatchEndpoint("https://order-service-623867952445.us-central1.run.app/orders/"+event.OrderID+"/status", &order)

		msg.Ack()
	})

	if err != nil {
		log.Fatalf("Error receiving messages: %v", err)
	}
}

func callPatchEndpoint(url string, payload interface{}) {
	// Convert payload to JSON
	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// Create a new PATCH request
	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTI3NjExNTIsImlhdCI6MTc1Mjc0Njc1MiwidXNlcm5hbWUiOiJzaGl2YW5pZ3VwdGEifQ.BEBMHZwx2fS3ztqjTHOj2LYV4JfEpZaKlDAaY74p9U8")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Status:", resp.Status)
}
