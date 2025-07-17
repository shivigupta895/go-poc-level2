package handlers

import (
	"net/http"
	"time"

	"order-service/models"
	"order-service/pubsub"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CreateOrder handles POST /orders
func CreateOrder(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var order models.Order
		if err := c.ShouldBindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		order.ID = uuid.New().String()
		order.Status = "PENDING"
		order.CreatedAt = time.Now()

		if err := db.Create(&order).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Publish order creation event
		pubsub.PublishOrderCreated(order)

		c.JSON(http.StatusCreated, order)
	}
}

// GetOrder handles GET /orders/:id
func GetOrder(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var order models.Order
		id := c.Param("id")

		if err := db.First(&order, "id = ?", id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
			return
		}

		c.JSON(http.StatusOK, order)
	}
}

// UpdateOrderStatus handles PATCH /orders/:id/status
func UpdateOrderStatus(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var input struct {
			Status string `json:"status"`
		}

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Model(&models.Order{}).Where("id = ?", id).Update("status", input.Status).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Order status updated"})
	}
}
