package routes

import (
	"order-service/config"
	"order-service/handlers"
	"order-service/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	db := config.InitDB()
	r := gin.Default()

	r.Use(middleware.JWTAuth())

	r.POST("/orders", handlers.CreateOrder(db))
	r.GET("/orders/:id", handlers.GetOrder(db))
	r.PATCH("/orders/:id/status", handlers.UpdateOrderStatus(db))

	return r
}
