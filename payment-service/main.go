package main

import (
	"log"
	"payment-service/config"
	"payment-service/pubsub"
)

func main() {
	db := config.InitDB()
	log.Println("Starting Payment Service...")

	pubsub.SubscribeToOrderEvents(db)
}
