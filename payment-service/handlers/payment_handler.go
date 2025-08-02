package handlers

import (
	"log"
	"net/http"
	"payment-service/config"
	"payment-service/pubsub"

	"github.com/gin-gonic/gin"
)

func PaymentHandler(c *gin.Context) {
	db := config.InitDB()
	log.Println(`{"message":"Inside payment handler", "service":"payment", "severity":"INFO"}`)

	pubsub.SubscribeToOrderEvents(db)
	c.JSON(http.StatusOK, gin.H{"message": "Payment Service started and subscribed to order events"})
}
