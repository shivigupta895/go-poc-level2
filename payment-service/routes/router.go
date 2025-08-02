package routes

import (
	"payment-service/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/payment", handlers.PaymentHandler)
	return r
}
