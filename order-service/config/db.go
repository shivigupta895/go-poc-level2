package config

import (
	"fmt"
	"log"
	"order-service/models"
	"order-service/utils"
	"os"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var GcpPojectId string = "go-poc-58"
var OrderTopicId string = "tp-order-events"
var OrderSubId string = "sb-order-events"

func InitDB() *gorm.DB {
	dbUser := os.Getenv("DB_USER")
	dbPassword, err := utils.GetSecret("DB_PASSWORD")
	if err != nil {
		log.Fatalf("Failed to get DB_PASSWORD: %v", err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	database := os.Getenv("DATABASE")

	conString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, database)
	log.Println("conString", conString)

	db, err := gorm.Open(mysql.Open(conString), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	db.AutoMigrate(&models.Order{})
	return db
}

func UpdateOrder(orderID string) (string, error) {
	db := InitDB()
	var order models.Order
	if err := db.First(&order, "id = ?", orderID).Error; err != nil {
		log.Printf(`{"message":"Order not found: %v", "service":"order", "severity":"ERROR"}`, err)
		return "", err
	}

	err := db.Model(&order).Where("id = ?", orderID).Update("status", "PAID").Error
	if err != nil {
		log.Printf(`{"message":"Failed to update order: %v", "service":"order", "severity":"ERROR"}`, err)
		return "", err
	}

	log.Printf(`{"message":"Order status updated to PAID for: %s", "service":"order", "severity":"INFO"}`, orderID)
	return "Order status updated to PAID", nil
}
