package main

import (
	"log"
	"net/http"
	"payment-service/config"
	"payment-service/pubsub"
)

func startPaymentService(w http.ResponseWriter, r *http.Request) {
	db := config.InitDB()
	log.Println("Starting Payment Service...")

	pubsub.SubscribeToOrderEvents(db)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Payment Service started and subscribed to order events."))
}

func main() {
	http.HandleFunc("/start-payment-service", startPaymentService)

	log.Println("Server is running on port 8081...")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
