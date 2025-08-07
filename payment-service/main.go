package main

import (
	"log"
	"payment-service/config"
	"payment-service/pubsub"
	"payment-service/routes"
	"payment-service/utils"
)

func main() {
	utils.LoadEnvVariables()
	db := config.InitDB()
	r := routes.SetupRouter()

	// Run the HTTP server in a goroutine
	go func() {
		if err := r.Run(":8081"); err != nil {
			log.Fatalf("Failed to run server: %v", err)
		}
	}()

	// Run the Pub/Sub subscriber in a goroutine
	go func() {
		pubsub.SubscribeToOrderEvents(db)
	}()

	// Block main from exiting
	select {}
}
