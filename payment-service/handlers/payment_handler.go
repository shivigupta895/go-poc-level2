package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PaymentHandler(c *gin.Context) {
	log.Println(`{"message":"Inside payment handler", "service":"payment", "severity":"INFO"}`)
	c.JSON(http.StatusOK, gin.H{"message": "Payment Service started and subscribed to order events"})
}
