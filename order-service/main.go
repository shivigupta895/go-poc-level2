package main

import (
	"log"
	"order-service/config"
	"order-service/handlers"
	"order-service/middleware"
	"order-service/pubsub"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.InitDB()
	r := gin.Default()

	r.Use(middleware.JWTAuth())

	r.POST("/orders", handlers.CreateOrder(db))
	r.GET("/orders/:id", handlers.GetOrder(db))
	r.PATCH("/orders/:id/status", handlers.UpdateOrderStatus(db))

	log.Println("Starting Order Service...")

	// Run the HTTP server in a goroutine
	go func() {
		if err := r.Run(":8080"); err != nil {
			log.Fatalf("Failed to run server: %v", err)
		}
	}()

	// Run the Pub/Sub subscriber in a goroutine
	go func() {
		pubsub.SubscribeToPaymentEvents(db)
	}()

	// Block main from exiting
	select {}
}
