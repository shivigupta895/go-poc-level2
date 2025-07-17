package config

import (
	"log"
	"order-service/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	//%0,5:F/tm@]tBD[d Password
	dsn := "shivani:%0,5:F/tm@]tBD[d@tcp(35.184.90.157:3306)/orders_db?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	db.AutoMigrate(&models.Order{})
	return db
}
