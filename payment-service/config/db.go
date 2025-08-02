package config

import (
	"fmt"
	"log"
	"os"
	"payment-service/models"
	"payment-service/utils"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var GcpPojectId string = "go-poc-58"
var PaymentTopicId string = "tp-payment-events"
var PaymentSubId string = "sb-payment-events"

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

	db.AutoMigrate(&models.Payment{})
	return db
}
